package media

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *MediaService) getComments(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "mediaid")
	id, _ := strconv.Atoi(param)

	query := `
		SELECT
			author,
			comment,
			TIMESTAMPDIFF(SECOND, post_date, NOW())
		FROM
			web_media_posts
		WHERE
			mediaid = ?
		ORDER BY
			postid
		DESC
	`
	result, err := s.db.Query(query, id)

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	comment := mediaComments{}
	comments := []mediaComments{}

	for result.Next() {
		result.Scan(&comment.Author, &comment.Comment, &comment.Date)
		comments = append(comments, comment)
	}
	result.Close()

	render.JSON(w, r, comments)
	return
}
