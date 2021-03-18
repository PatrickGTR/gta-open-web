package media

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

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
