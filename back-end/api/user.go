package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/open-backend/user"
	"github.com/open-backend/util"
)

// VerifyUser (/user - POST)
func VerifyUser(w http.ResponseWriter, r *http.Request) {
	// grab data from the body (form-data)
	formUsername := r.FormValue("username")
	formPassword := r.FormValue("password")

	// retrieve password from database
	password := getPasswordFromDB(formUsername)
	match := util.ComparePassword(password, formPassword)
	if password == "" || !match {
		data := &Exception{
			Code:    "login.wrong.password",
			Message: "Oops something went wrong, try again.",
		}
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, data)

	} else {

		uid, err := getUserID(formUsername)
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

// Logout (/user - DELETE)
func Logout(w http.ResponseWriter, r *http.Request) {
	user.DestroySession(w, r)
	render.Status(r, http.StatusOK)
}

// GetDataByUID (/user/userid - GET)
// grabs all the user data that will be shown in dashboard.
// or other parts of the website.
func GetDataByUID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "userid")
	userid, _ := strconv.Atoi(param)
	data, err := getAllData(userid)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	render.JSON(w, r, &Player{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items})
	render.Status(r, http.StatusOK)
}
