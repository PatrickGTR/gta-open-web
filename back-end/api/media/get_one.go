package media

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

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
