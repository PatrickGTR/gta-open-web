package session

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
)

var Session *sessions.CookieStore

func New() *sessions.CookieStore {

	Session = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	Session.Options.Path = "/"
	Session.Options.HttpOnly = true

	state := false

	if os.Getenv("ENV") != "DEV" {
		state = true
		Session.Options.Domain = "gta-open.ga"
	}
	Session.Options.Secure = state

	return Session
}

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

// Destroy deletes and destroys the cookie token that are stored in users
// browser and deletes the data on that token.
func Destroy(w http.ResponseWriter, r *http.Request) {
	session, _ := Session.Get(r, "sessionid")
	session.Options.MaxAge = -1
	session.Save(r, w)

}

func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// grab the content of the session.
		// if there are none set, return unauthorized http status
		_, err := GetUID(r)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			return
		}

		// proceed to the next route
		next(w, r)
	})
}
