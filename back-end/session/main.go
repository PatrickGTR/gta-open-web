package session

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Session *sessions.CookieStore

func init() {

	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	Session = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	Session.Options.Path = "/"
	Session.Options.HttpOnly = true

	state := false

	if os.Getenv("ENV") != "DEV" {
		state = true
		Session.Options.Domain = "gta-open.ga"
	}
	Session.Options.Secure = state

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
