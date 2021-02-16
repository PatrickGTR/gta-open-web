package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func MediaGetAll(w http.ResponseWriter, r *http.Request) {

	query := `
		SELECT
			link,
			author,
			post_date,
			views
		FROM
			web_media
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
		result.Scan(&post.Link, &post.Author, &post.Time, &post.Views)

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

	err = insertMediaToDB(body.Link, body.Author)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: "Could write to database.",
		})
		fmt.Println(err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &Exception{
		Code:    "media.posted",
		Message: "You have succesfully posted your content",
	})
	return
}
