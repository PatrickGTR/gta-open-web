package user

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

var Cookie *sessions.CookieStore

func init() {
	Cookie = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	Cookie.Options.Path = "/"
	Cookie.Options.HttpOnly = true
}

func GenerateSession(w http.ResponseWriter, r *http.Request, uid int) (err error) {
	// 90 days expiration
	expiration := time.Now().Add((time.Hour * 24) * 3)
	http.SetCookie(w, &http.Cookie{
		Name:    "db_user_id",
		Value:   strconv.Itoa(uid),
		Path:    "/",
		Expires: expiration,
	})

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
	http.SetCookie(w, &http.Cookie{
		Name:   "db_user_id",
		Path:   "/",
		MaxAge: -1,
	})

	session, _ := Cookie.Get(r, "sessionid")
	session.Options.MaxAge = -1
	session.Save(r, w)

	return
}
