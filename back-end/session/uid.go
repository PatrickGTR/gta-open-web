package session

import (
	"errors"
	"net/http"
)

// GetUID retrieves the account unique id based on the session token.
// this function returns the unique id and an error
func GetUID(r *http.Request) (uid int, err error) {
	session, _ := Session.Get(r, "sessionid")
	sessionContent := session.Values["accountID"]
	if sessionContent == nil {
		err = errors.New("no session set")
		return
	}

	uid = session.Values["accountID"].(int)
	if uid <= 0 {
		err = errors.New("invalid session userid <= 0")
		return
	}
	return
}
