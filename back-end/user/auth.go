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

func GetData(uid int) (data Player, err error) {
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

	result, err := helper.ExecuteQuery(query, uid)
	if err != nil {
		return
	}

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

		return
	}
	return
}
