package user

import (
	"fmt"

	"github.com/open-backend/helper"
)

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
	return
}
