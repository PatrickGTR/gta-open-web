package server

import (
	"errors"

	"github.com/open-backend/helper"
)

func GetHighest(highestType uint8) (username string, err error) {
	option := ""

	switch highestType {
	case HIGHEST_KILLS:
		option = "ps.kills"
	case HIGHEST_DEATHS:
		option = "ps.deaths"
	case HIGHEST_MONEY:
		option = "ps.money"
	default:
		err = errors.New("Invalid type for function 'GetHighest'")
		return
	}

	query := `
		SELECT
			p.username
		FROM
			player_stats ps
		INNER JOIN
			players p
		ON
			p.u_id = ps.u_id
		ORDER BY
	`
	query = query + " " + option + " DESC LIMIT 1"
	result, err := helper.ExecuteQuery(query)
	if err != nil {
		return
	}

	result.Next()
	result.Scan(&username)
	result.Close()
	return
}
