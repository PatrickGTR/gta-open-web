package media

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

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
