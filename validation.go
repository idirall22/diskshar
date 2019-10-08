package user

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorEmailNotValid    = errors.New("Email not valid")
	ErrorUsernameNotValid = errors.New("Username not valid")
)

// validate email
func validateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	valid, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if !valid {
		return email, ErrorEmailNotValid
	}
	return email, nil
}

// validate username
func validateUsername(username string) (string, error) {
	username = strings.TrimSpace(username)
	valid, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username)

	if !valid {
		return username, ErrorUsernameNotValid
	}
	return username, nil
}
