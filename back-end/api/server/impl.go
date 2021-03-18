package server

import (
	"errors"
	"fmt"

	"github.com/open-backend/util"
)

// Bans Structure
type Bans struct {
	Username  string `json:"username"`
	BannedBy  string `json:"by"`
	Reason    string `json:"reason"`
	BanDate   string `json:"banDate"`
	UnbanDate string `json:"unbanDate"`
}

type Exception util.MessageData

const (
	highestKills  = 1
	highestMoney  = 2
	highestDeaths = 3
)
const (
	dataHighest      = 1
	dataTotalAccount = 2
)

func (s *ServerService) getServerData(arg ...uint8) (data interface{}, err error) {

	dataType := arg[0]
	highestType := arg[1]

	switch dataType {
	case dataHighest:
		data, err = s.getHighest(highestType)
		return
	case dataTotalAccount:
		data, err = s.getRegisteredPlayers()
		return
	default:
		err = errors.New("Invalid data type")
		return
	}
}

func (s *ServerService) getRegisteredPlayers() (total int, err error) {
	query := `
		SELECT
			COUNT(*)
		FROM
			players
	`

	result, err := s.db.Query(query)
	result.Next()
	result.Scan(&total)
	result.Close()
	return
}

func (s *ServerService) getHighest(highestType uint8) (username string, err error) {
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
	result, err := s.db.Query(query)

	result.Next()
	result.Scan(&username)
	result.Close()
	return
}

func (s *ServerService) getAllBans() ([]Bans, error) {

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
	result, err := s.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var bans = Bans{}
	banArr := []Bans{}
	for result.Next() {
		result.Scan(
			&bans.Username,
			&bans.BannedBy,
			&bans.Reason,
			&bans.BanDate,
			&bans.UnbanDate)

		banArr = append(banArr, bans)
	}
	result.Close()
	return banArr, nil
}
