package session

import "net/http"

// Generate creates a new session when this function gets called, it stores
// the user's unique ID & admin level
func Generate(w http.ResponseWriter, r *http.Request, uid int) (err error) {

	session, _ := Session.Get(r, "sessionid")
	session.Values["accountID"] = uid
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		return
	}
	return
}
