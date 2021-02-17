package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/open-backend/user"
)

func MediaGetOne(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	fmt.Println(id)

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

	result, err := ExecuteQuery(query, id)
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

func MediaGetAll(w http.ResponseWriter, r *http.Request) {
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
	result, err := ExecuteQuery(query)
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
func MediaPost(w http.ResponseWriter, r *http.Request) {
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
	userid, _ := user.GetUIDFromSession(r)
	result, _ := ExecuteQuery("SELECT username FROM players WHERE u_id = ?", userid)
	result.Next()
	result.Scan(&body.Author)
	result.Close()

	err = insertMediaToDB(body.Link, body.Title, body.Author)
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
