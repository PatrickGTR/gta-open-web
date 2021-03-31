package session

import "net/http"

// Destroy deletes and destroys the cookie token that are stored in users
// browser and deletes the data on that token.
func Destroy(w http.ResponseWriter, r *http.Request) {
	session, _ := Session.Get(r, "sessionid")
	session.Options.MaxAge = -1
	session.Save(r, w)
}
