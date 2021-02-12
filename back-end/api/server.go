package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

// ServerStats (/server/stats)
func ServerStats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	dataType, _ := strconv.Atoi(params.Get("type"))
	option, _ := strconv.Atoi(params.Get("option"))
	data, err := getServerData(uint8(dataType), uint8(option))

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "invalid.type",
			Message: "Invalid type value",
		})
		return
	}

	render.JSON(w, r, struct {
		Data interface{} `json:"value"`
	}{
		Data: data,
	})
	render.Status(r, http.StatusOK)
	return
}

func BanList(w http.ResponseWriter, r *http.Request) {

	query := `
		SELECT
			username,
			admin,
			reason,
			DATE_FORMAT(ban_date, "%d %b %Y at %h:%i %p"),
			DATE_FORMAT(unban_date, "%d %b %Y at %h:%i %p")
		FROM
			bans
	`
	result, err := ExecuteQuery(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var bans = Bans{}
	returnBans := []Bans{}
	for result.Next() {
		result.Scan(
			&bans.Username,
			&bans.BannedBy,
			&bans.Reason,
			&bans.BanDate,
			&bans.UnbanDate)

		returnBans = append(returnBans, bans)
	}
	result.Close()

	render.JSON(w, r, returnBans)
	render.Status(r, http.StatusOK)
	return
}
