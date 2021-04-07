package media

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

func (s *MediaService) deleteMedia(w http.ResponseWriter, r *http.Request) {
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

	sessionUID, err := session.GetUID(r)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: err.Error(),
		})
		return
	}

	query := `
	SELECT
		p.u_id
	FROM
		players p
	INNER JOIN
		web_media w
	ON
		p.username = w.author
	WHERE
		w.author = ?
	AND
		w.postid = ?
	`
	result := s.db.QueryRow(query, body.Author, body.Postid)
	posterUID := 0
	result.Scan(&posterUID)

	if sessionUID != posterUID {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, &Exception{
			Code:    "not.poster",
			Message: "session uid not the same as db uid",
		})
		return
	}

	query = `
	DELETE FROM
		web_media
	WHERE
		postid = ?
	`
	_, err = s.db.Exec(query, body.Postid)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &Exception{
		Code:    "post.deleted",
		Message: fmt.Sprintf("Successfully deleted postid %d", body.Postid),
	})
	return
}
