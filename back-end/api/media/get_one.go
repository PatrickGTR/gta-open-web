package media

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// grabs the post data (link, title, author, date posted and views) ...
// by the given post id.
func (s *MediaService) getOne(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	post, err := s.getMediaFromDB(id)
	if err != nil {
		render.JSON(w, r, &Exception{
			Code:    "internal.error",
			Message: "Issue fetching data, check database connection.",
		})
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, post)
	render.Status(r, http.StatusOK)
	return
}
