package media

import (
	"database/sql"
	"net/http"
	"strconv"

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
