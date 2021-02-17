package api

type mediaBody struct {
	Postid int    `json:"id"`
	Link   string `json:"youtubeLink"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Time   string `json:"datePosted"`
	Views  int    `json:"views"`
}

func insertMediaToDB(link string, title string, author string) (err error) {
	query := `INSERT INTO web_media (link, title, author) VALUES (?, ?, ?)`
	_, err = ExecuteQuery(query, link, title, author)
	return
}
