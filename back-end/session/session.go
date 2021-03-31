package session

import (
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
