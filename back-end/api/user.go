package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/open-backend/session"
	"github.com/open-backend/util"
)

// VerifyUser (/user - POST)
// HTTP RESPONSES
// 200 -- Success, username and password is good.
// 400 -- Bad Request, json is invalid or username, password is empty.
// 401 -- Unauthorized, username is valid but user input password does not match
// ... the password in the database.
func VerifyUser(w http.ResponseWriter, r *http.Request) {
	bodyResponse := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // catch unwanted fields
	err := decoder.Decode(&bodyResponse)

	if err != nil {
		// bad json data has been passed
		// could be unrecognized field.
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "bad.data.passed",
			Message: err.Error(),
		})
		return
	}

	username := bodyResponse.Username
	password := bodyResponse.Password
	if username == "" || password == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "missing.data.field",
			Message: "username or password is empty, required",
		})
		return
	}

	dbPassword := getPasswordFromDB(username)
	match := util.ComparePassword(dbPassword, password)
	if password == "" || !match {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, &Exception{
			Code:    "login.wrong.password",
			Message: "Oops something went wrong, try again.",
		})
		return
	}

	// if error occurs below, just print it out to the console
	// the client does not need to see why as it's mostly internal error.
	// but can be good for devs to catch wrong queries etc..
	// revise this code, it probably does need any error handling tbh.
	uid, err := getUserID(username)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = session.Generate(w, r, uid)
	if err != nil {
		fmt.Println(err.Error())
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &Exception{
		Code:    "login.success",
		Message: "You have successfully logged in",
	})
	return
}

// Logout (/user - DELETE)
func Logout(w http.ResponseWriter, r *http.Request) {
	session.Destroy(w, r)
	render.Status(r, http.StatusOK)
	return
}

func GetSelfData(w http.ResponseWriter, r *http.Request) {
	userid, _ := session.GetUID(r)
	data, err := getAllData(userid)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	render.JSON(w, r, &Player{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items})
	render.Status(r, http.StatusOK)
	return
}

// GetDataByUID (/user/userid - GET)
// grabs all the user data that will be shown in dashboard.
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
	return
}
