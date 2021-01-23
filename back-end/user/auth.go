package user

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/open-backend/helper"
)

type JWTData struct {
	UID   int   `json:"uid"`
	Admin uint8 `json:"admin"`
	jwt.StandardClaims
}

func GenerateToken(uid int, admin uint8) (signToken string, err error) {

	// expire in 30 days time
	exp := time.Now().Add((time.Hour * 24) * 30).Unix()

	claims := JWTData{
		UID:   uid,
		Admin: admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "GTA-Open",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signToken, err = token.SignedString([]byte("D887FF9A47539CF9592E5D14BEE93"))
	return
}

func GetUidAdmin(username string) (uid int, admin uint8) {

	query := `
		SELECT
			p.u_id,
			a.admin_level
		FROM
			players p
		LEFT JOIN
			admins a
		ON
			p.u_id = a.u_id
		WHERE
			p.username = ?
	`

	result, err := helper.ExecuteQuery(query, username)
	if err != nil {
		fmt.Println(err.Error())
	}

	result.Next()
	result.Scan(&uid, &admin)
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
	return
}
