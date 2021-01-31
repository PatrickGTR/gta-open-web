package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/open-backend/helper"
)

var Cookie *sessions.CookieStore

func init() {
	Cookie = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	Cookie.Options.Path = "/"
}

func GenerateSession(w http.ResponseWriter, r *http.Request, uid int) (err error) {

	session, _ := Cookie.Get(r, "sessionid")

	session.Values["accountID"] = uid
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "db_user_id",
		Value:  fmt.Sprint(uid),
		Path:   "/",
		MaxAge: 86400 * 30,
	})

	return
}

func GetUID(username string) (uid int, err error) {

	query := `
		SELECT
			u_id
		FROM
			players
		WHERE
			username = ?
	`

	result, err := helper.ExecuteQuery(query, username)
	if err != nil {
		return
	}

	result.Next()
	result.Scan(&uid)
	result.Close()
	return
}

func GetPassword(username string) (password string) {
	query := `
		SELECT
			password
		FROM
			players
		WHERE
			username = ?
	`

	result, err := helper.ExecuteQuery(query, username)
	if err != nil {
		fmt.Println(err.Error())
	}

	// grab single data
	result.Next()
	result.Scan(&password)
	result.Close()
	return
}
