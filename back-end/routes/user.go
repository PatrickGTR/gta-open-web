package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/open-backend/helper"
	"github.com/open-backend/user"
)

// Exception alias of helper.MessageData
type Exception helper.MessageData

func Login(w http.ResponseWriter, r *http.Request) {
	// grab data from the body (form-data)
	formUsername := r.FormValue("username")
	formPassword := r.FormValue("password")

	// retrieve password from database
	password := user.GetPassword(formUsername)
	match := helper.ComparePassword(password, formPassword)
	if password == "" || !match {
		data := &Exception{
			Code:    "login.wrong.password",
			Message: "Oops something went wrong, try again.",
		}
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, data)

	} else {

		uid, err := user.GetUID(formUsername)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = user.GenerateSession(w, r, uid)
		if err != nil {
			fmt.Println(err.Error())
		}

		data := &Exception{
			Code:    "login.success",
			Message: "You have successfully logged in",
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, data)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := user.Cookie.Get(r, "sessionid")

	// no session id set
	sessionContent := session.Values["accountID"]
	if sessionContent == nil {
		render.Status(r, http.StatusUnauthorized)
		return
	}

	// account ID starts at 1
	sessionUID := session.Values["accountID"].(int)
	if sessionUID <= 0 {
		render.Status(r, http.StatusUnauthorized)
		return
	}

	session.Options.MaxAge = -1
	session.Values["accountID"] = 0
	session.Save(r, w)

	render.Status(r, http.StatusOK)
}

func GetDataByUID(w http.ResponseWriter, r *http.Request) {

	session, _ := user.Cookie.Get(r, "sessionid")
	userid := session.Values["accountID"].(int)

	data, err := user.GetData(userid)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	render.JSON(w, r, &user.Player{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items})
	render.Status(r, http.StatusOK)
}
