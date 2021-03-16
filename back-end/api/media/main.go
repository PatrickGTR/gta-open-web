package media

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

type MediaService struct {
	Routes chi.Router
	db     *sql.DB
}

func New(db *sql.DB) *MediaService {
	router := chi.NewRouter()

	service := &MediaService{
		Routes: router,
		db:     db,
	}

	router.Post("/", session.WithAuthentication(service.addMedia))
	router.Get("/", service.getAll)
	router.Get("/{id}", service.getOne)
	router.Get("/trending", service.getTrending)
	router.Post("/add_views", service.incrementViews)

	router.Route("/comment", func(r chi.Router) {
		r.Post("/", session.WithAuthentication(service.postComment))
		r.Get("/{mediaid}", service.getComments)
	})

	return service
}

func (s *MediaService) incrementViews(w http.ResponseWriter, r *http.Request) {
	bodyResponse := struct {
		MediaID string `json:"mediaid"`
	}{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // catch unwanted fields
	err := decoder.Decode(&bodyResponse)

	if err != nil {
		// bad json data has been passed
		// could be unrecognized field.
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "bad.data.passed",
			Message: err.Error(),
		})
		return
	}

	query := `
		UPDATE
			web_media
		SET
			views = views + 1
		WHERE
			postid = ?
	`
	s.db.Query(query, bodyResponse.MediaID)
	return
}

func (s *MediaService) postComment(w http.ResponseWriter, r *http.Request) {
	bodyResponse := struct {
		MediaID string `json:"mediaid"`
		Comment string `json:"comment"`
	}{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // catch unwanted fields
	err := decoder.Decode(&bodyResponse)

	if err != nil {
		// bad json data has been passed
		// could be unrecognized field.
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "bad.data.passed",
			Message: err.Error(),
		})
		return
	}
	// get the name based on account id
	var author string
	userid, _ := session.GetUID(r)
	result, _ := s.db.Query("SELECT username FROM players WHERE u_id = ?", userid)
	result.Next()
	result.Scan(&author)
	result.Close()

	query :=
		`
		INSERT INTO
			web_media_posts (mediaid, author, comment)
		VALUES
			(?, ?, ?)
	`
	_, err = s.db.Query(query, bodyResponse.MediaID, author, bodyResponse.Comment)
	return
}

func (s *MediaService) getComments(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "mediaid")
	id, _ := strconv.Atoi(param)

	query := `
		SELECT
			author,
			comment,
			TIMESTAMPDIFF(SECOND, post_date, NOW())
		FROM
			web_media_posts
		WHERE
			mediaid = ?
		ORDER BY
			postid
		DESC
	`
	result, err := s.db.Query(query, id)

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	comment := mediaComments{}
	comments := []mediaComments{}

	for result.Next() {
		result.Scan(&comment.Author, &comment.Comment, &comment.Date)
		comments = append(comments, comment)
	}
	result.Close()

	render.JSON(w, r, comments)
	return
}

func (s *MediaService) getTrending(w http.ResponseWriter, r *http.Request) {
	webQuery := r.URL.Query().Get("q")

	post := mediaBody{}

	switch webQuery {
	case "hottest":
		{
			result, _ := s.db.Query(`
				SELECT
					title,
					link,
					author,
					views
				FROM
					web_media
				ORDER BY
					views
				DESC LIMIT 1
			`)

			result.Next()
			result.Scan(&post.Title, &post.Link, &post.Author, &post.Views)
			result.Close()

		}
	case "newest":
		{
			result, _ := s.db.Query(`
				SELECT
					title,
					link,
					author,
					views
				FROM
					web_media
				ORDER BY
					postid
				DESC LIMIT 1
			`)

			result.Next()
			result.Scan(&post.Title, &post.Link, &post.Author, &post.Views)
			result.Close()

		}
	default:
		render.JSON(w, r, &Exception{
			Code:    "invalid.query.value",
			Message: "Acceptable queries are 'hotest', 'newest'",
		})
		render.Status(r, http.StatusNotFound)
		return
	}

	render.JSON(w, r, post)
	render.Status(r, http.StatusOK)
	return
}

func (s *MediaService) getOne(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	query := `
	SELECT
			link,
			title,
			author,
			DATE_FORMAT(post_date, "%d %M %Y at %h:%i%p"),
			views
		FROM
			web_media
		WHERE
			postid = ?`

	result, err := s.db.Query(query, id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}

	post := mediaBody{}

	result.Next()
	result.Scan(&post.Link, &post.Title, &post.Author, &post.Time, &post.Views)
	result.Close()

	render.JSON(w, r, post)
	render.Status(r, http.StatusOK)
	return
}

func (s *MediaService) getAll(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			postid,
			link,
			title,
			author,
			TIMESTAMPDIFF(SECOND, post_date, NOW()),
			views
		FROM
			web_media
		ORDER BY
			postid
		DESC
	`
	result, err := s.db.Query(query)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: "Could not retrieve data",
		})
		return
	}

	post := mediaBody{}
	allPost := []mediaBody{}

	for result.Next() {
		result.Scan(&post.Postid, &post.Link, &post.Title, &post.Author, &post.Time, &post.Views)

		allPost = append(allPost, post)
	}
	result.Close()

	render.JSON(w, r, allPost)
	return
}

// MediaPost - HTTP Responses
// 200 - success
// 400 - invalid json passed
// 406 - invalid link
// 500 - most likely mysql error.
func (s *MediaService) addMedia(w http.ResponseWriter, r *http.Request) {
	var err error
	body := mediaBody{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // catch unwanted fields
	err = decoder.Decode(&body)
	if err != nil {
		// bad json data has been passed
		// could be unrecognized field.
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "bad.data.passed",
			Message: err.Error(),
		})
		return
	}

	isYoutube := strings.Contains(body.Link, "www.youtube.com")
	if !isYoutube {
		render.Status(r, http.StatusNotAcceptable)
		render.JSON(w, r, &Exception{
			Code:    "invalid.link",
			Message: "Only youtube links allowed",
		})
		return
	}

	titleLength := len(body.Title)
	if titleLength > 50 {
		render.Status(r, http.StatusNotAcceptable)
		render.JSON(w, r, &Exception{
			Code:    "title.too.long",
			Message: "The maximum characters of title is 50",
		})
	}

	// get the name based on account id
	userid, _ := session.GetUID(r)
	result, _ := s.db.Query("SELECT username FROM players WHERE u_id = ?", userid)
	result.Next()
	result.Scan(&body.Author)
	result.Close()

	err = s.insertMediaToDB(body.Link, body.Title, body.Author)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: "Could not write to database.",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &Exception{
		Code:    "media.posted",
		Message: "You have succesfully posted your content",
	})
	return
}
