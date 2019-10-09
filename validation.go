package user

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrorEmailNotValid when email is not valid
	ErrorEmailNotValid = errors.New("Email not valid")

	// ErrorUsernameNotValid when username is not valid
	ErrorUsernameNotValid = errors.New("Username not valid")
)

// validate email
func validateEmail(email string) (string, error) {

	if email == "" {
		return email, ErrorEmailNotValid
	}

	email = strings.TrimSpace(email)
	valid, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if !valid {
		return email, ErrorEmailNotValid
	}
	return email, nil
}

// validate username
func validateUsername(username string) (string, error) {

	if username == "" {
		return username, ErrorUsernameNotValid
	}
	username = strings.TrimSpace(username)
	valid, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username)

	if !valid {
		return username, ErrorUsernameNotValid
	}
	return username, nil
}

func parseUniqueConstraintError(err string) error {

	e := strings.Split(err, "_")[1]
	switch e {
	case "email":
		return ErrorEmailTaken
	case "username":
		return ErrorUsernameTaken
	}
	return nil
}

// validateLoginForm
func validateLoginForm(form *RegisterForm) (*RegisterForm, error) {

	// Validate username
	username, errU := validateUsername(form.Username)
	if errU != nil {
		return nil, errU
	}
	form.Username = username

	// Validate email
	email, errE := validateEmail(form.Email)
	if errE != nil {
		return nil, errE
	}
	form.Email = email

	return form, nil
}
