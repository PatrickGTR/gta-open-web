package media

import (
	"github.com/open-backend/util"
)

type mediaBody struct {
	Postid int    `json:"id"`
	Link   string `json:"youtubeLink"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Time   string `json:"datePosted"`
	Views  int    `json:"views"`
}

type mediaComments struct {
	Author  string `json:"author"`
	Comment string `json:"comment"`
	Date    string `json:"datePosted"`
}

type Exception util.MessageData

func (s *MediaService) insertMediaToDB(link string, title string, author string) (err error) {
	query := `INSERT INTO web_media (link, title, author) VALUES (?, ?, ?)`
	_, err = s.db.Query(query, link, title, author)
	return
}

func (s *MediaService) getAllMediaFromDB() ([]mediaBody, error) {
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
		return []mediaBody{}, err
	}

	post := mediaBody{}
	allPost := []mediaBody{}

	for result.Next() {
		result.Scan(&post.Postid, &post.Link, &post.Title, &post.Author, &post.Time, &post.Views)
		allPost = append(allPost, post)
	}
	result.Close()
	return allPost, nil
}

func (s *MediaService) getMediaFromDB(postid int) (mediaBody, error) {
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

	result, err := s.db.Query(query, postid)
	if err != nil {
		return mediaBody{}, err
	}

	post := mediaBody{}

	result.Next()
	result.Scan(&post.Link, &post.Title, &post.Author, &post.Time, &post.Views)
	result.Close()
	return post, nil
}
