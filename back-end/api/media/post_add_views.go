package media

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

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
