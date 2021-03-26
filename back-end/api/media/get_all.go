package media

import (
	"net/http"

	"github.com/go-chi/render"
)

// grabs all the media's postid, link, title, author...
// date posted and total views from the database
func (s *MediaService) getAll(w http.ResponseWriter, r *http.Request) {

	data, err := s.getAllMediaFromDB()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: "Could not retrieve data",
		})
	}

	render.JSON(w, r, data)
	return
}
