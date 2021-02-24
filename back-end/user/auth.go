package user

import (
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Cookie *sessions.CookieStore

func init() {
	Cookie = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	Cookie.Options.Path = "/"
	Cookie.Options.HttpOnly = true
	Cookie.Options.SameSite = http.SameSiteNoneMode

	state := false
	if os.Getenv("ENV") == "PROD" {
		state = true
	}
	Cookie.Options.Secure = state

}

func GetUIDFromSession(r *http.Request) (uid int, err error) {
	session, _ := Cookie.Get(r, "sessionid")
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

func GenerateSession(w http.ResponseWriter, r *http.Request, uid int) (err error) {

	session, _ := Cookie.Get(r, "sessionid")
	session.Values["accountID"] = uid
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		return
	}
	return
}

func DestroySession(w http.ResponseWriter, r *http.Request) {

	session, _ := Cookie.Get(r, "sessionid")
	session.Options.MaxAge = -1
	session.Save(r, w)

	return
}
