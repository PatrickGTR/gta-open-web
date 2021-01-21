package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/open-backend/helper"
	"github.com/open-backend/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// grab data from the body (form-data)
	username := r.FormValue("username")
	formPassword := r.FormValue("password")

	// retrieve password from database
	password := user.GetPassword(username)
	match := helper.ComparePassword(password, formPassword)

	data := []helper.MessageData{}
	if password == "" || !match {
		data = helper.SendMessage("login.wrong.password", "Oops something went wrong, try again.")
		render.Status(r, http.StatusUnauthorized)

	} else {

		data = helper.SendMessage("login.success", "You have succesfully logged in.")
		render.Status(r, http.StatusOK)
	}

	render.JSON(w, r, data)
}

func GetDataByUID(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")

	query := `
		SELECT
			p.u_id,
			p.username,
			p.register_date,
			p.last_login,
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

		data := helper.SendMessage("mysql.error", fmt.Sprintf("%s", err))

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, data)
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
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, helper.SendMessage("no.rows", "Oops something went wrong, contact developer ASAP."))
		return
	}

	render.JSON(w, r, []user.Player{{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items}})
}