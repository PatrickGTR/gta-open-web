package util

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(passwordFirst string, passwordSecond string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordFirst), []byte(passwordSecond))
	if err != nil {
		return false
	}
	return true
}
