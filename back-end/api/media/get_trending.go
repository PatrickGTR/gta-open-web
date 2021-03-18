package media

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s *MediaService) getTrending(w http.ResponseWriter, r *http.Request) {
	webQuery := r.URL.Query().Get("q")

	post := mediaBody{}

	switch webQuery {
	case "hottest":
		{
			result, _ := s.db.Query(`
				SELECT
					title,
					link,
					author,
					views
				FROM
					web_media
				ORDER BY
					views
				DESC LIMIT 1
			`)

			result.Next()
			result.Scan(&post.Title, &post.Link, &post.Author, &post.Views)
			result.Close()

		}
	case "newest":
		{
			result, _ := s.db.Query(`
				SELECT
					title,
					link,
					author,
					views
				FROM
					web_media
				ORDER BY
					postid
				DESC LIMIT 1
			`)

			result.Next()
			result.Scan(&post.Title, &post.Link, &post.Author, &post.Views)
			result.Close()

		}
	default:
		render.JSON(w, r, &Exception{
			Code:    "invalid.query.value",
			Message: "Acceptable queries are 'hotest', 'newest'",
		})
		render.Status(r, http.StatusNotFound)
		return
	}

	render.JSON(w, r, post)
	render.Status(r, http.StatusOK)
	return
}
