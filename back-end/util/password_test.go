package util

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestWithoutMatchingPassword(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("hello"), bcrypt.DefaultCost)

	match := ComparePassword(string(hashedPassword), "world")
	if match == true {
		t.Error("Result was incorrect, got:", match, "want:", !match)
	}
}

func TestWithMatchingPassword(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("hello"), bcrypt.DefaultCost)

	match := ComparePassword(string(hashedPassword), "hello")
	if match == false {
		t.Error("password does not match")
	}
}
