package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

	http.SetCookie(w, &http.Cookie{
		Name:   "db_user_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	render.Status(r, http.StatusOK)
}

func GetDataByUID(w http.ResponseWriter, r *http.Request) {

	userid, _ := strconv.Atoi(chi.URLParam(r, "userid"))

	query := `
		SELECT
			p.u_id,
			p.username,
			p.register_date,
			IFNULL(p.last_login, "N/A"),
			ps.kills,
			ps.deaths,
			ps.money,
			ps.job_id,
			ps.class_id,
			ps.score,
			ps.skin,
			IFNULL(pi.crack, 0),
			IFNULL(pi.weed, 0),
			IFNULL(pi.picklock, 0),
			IFNULL(pi.wallet, 0),
			IFNULL(pi.rope, 0),
			IFNULL(pi.condom, 0),
			IFNULL(pi.scissors, 0)
		FROM
			players p
		INNER JOIN
			player_stats ps
		ON
			p.u_id = ps.u_id
		LEFT JOIN
			player_items pi
		ON
			p.u_id = pi.u_id
		WHERE
			p.u_id = ?
		`

	result, err := helper.ExecuteQuery(query, userid)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "mysql.error",
			Message: fmt.Sprintf("%s", err),
		})
		return
	}

	var data user.Player
	// store data
	result.Next()
	err = result.Scan(
		&data.Account.UID,
		&data.Account.Username,
		&data.Account.Registered,
		&data.Account.LastLogin,
		&data.Stats.Kills,
		&data.Stats.Deaths,
		&data.Stats.Money,
		&data.Stats.Job,
		&data.Stats.Class,
		&data.Stats.Score,
		&data.Stats.Skin,
		&data.Items.Crack,
		&data.Items.Weed,
		&data.Items.Picklock,
		&data.Items.Wallet,
		&data.Items.Rope,
		&data.Items.Condom,
		&data.Items.Scissors,
	)
	result.Close()
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "no.rows",
			Message: err.Error()},
		)
		return
	}

	render.JSON(w, r, &user.Player{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items})
}
