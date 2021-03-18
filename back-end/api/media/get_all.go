package media

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s *MediaService) getAll(w http.ResponseWriter, r *http.Request) {
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
	result, err := s.db.Query(query)
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
