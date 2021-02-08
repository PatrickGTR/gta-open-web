package api

import (
	"errors"
	"fmt"
)

func getServerData(arg ...uint8) (data interface{}, err error) {

	dataType := arg[0]
	highestType := arg[1]

	switch dataType {
	case dataHighest:
		data, err = getHighest(highestType)
		return
	case dataTotalAccount:
		data, err = getRegisteredPlayers()
		return
	default:
		err = errors.New("Invalid data type")
		return
	}
}

func getRegisteredPlayers() (total int, err error) {
	query := `
		SELECT
			COUNT(*)
		FROM
			players
	`

	result, err := ExecuteQuery(query)
	result.Next()
	result.Scan(&total)
	result.Close()
	return
}

func getHighest(highestType uint8) (username string, err error) {
	option := ""

	switch highestType {
	case highestKills:
		option = "ps.kills"
	case highestDeaths:
		option = "ps.deaths"
	case highestMoney:
		option = "ps.money"
	default:
		err = errors.New("Invalid type for function 'GetServerData'")
		return
	}

	query := fmt.Sprintf(`
		SELECT
			p.username
		FROM
			player_stats ps
		INNER JOIN
			players p
		ON
			p.u_id = ps.u_id
		ORDER BY
			%s
		DESC LIMIT 1
	`, option)
	result, err := ExecuteQuery(query)

	result.Next()
	result.Scan(&username)
	result.Close()
	return
}
