package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(passwordFirst string, passwordSecond string) (match bool) {

	err := bcrypt.CompareHashAndPassword([]byte(passwordFirst), []byte(passwordSecond))
	if err != nil {
		match = false
	} else {
		match = true
	}
	return
}
