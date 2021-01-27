package user

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/open-backend/helper"
)

type JWTData struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid int) (signToken string, err error) {

	// expire every 30 minute
	exp := time.Now().Add(time.Minute * 30).Unix()

	claims := JWTData{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "GTA-Open",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signToken, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
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
