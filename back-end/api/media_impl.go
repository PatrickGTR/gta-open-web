package api

type mediaBody struct {
	Link   string `json:"youtubeLink"`
	Author string `json:"author"`
	Time   string `json:"datePosted"`
	Views  int    `json:"views"`
}

func insertMediaToDB(link string, author string) (err error) {
	query := `INSERT INTO web_media (link, author) VALUES (?, ?)`
	_, err = ExecuteQuery(query, link, author)
	return
}
