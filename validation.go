package userAccount

import (
	"errors"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrorEmailNotValid when email is not valid
	ErrorEmailNotValid = errors.New("Email not valid")

	// ErrorUsernameNotValid when username is not valid
	ErrorUsernameNotValid = errors.New("Username not valid")

	// ErrorLoginInfos when username and email are not valid
	ErrorLoginInfos = errors.New("You need to use email or username to login")

	// ErrorTokenNoyValid when token is not valid
	ErrorTokenNoyValid = errors.New("token not valid")
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
func validateRegisterForm(form *RegisterForm) (*RegisterForm, error) {

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

	// validate password
	err := validatePassword(form.Password)

	if err != nil {
		return nil, err
	}

	return form, nil
}

// validateLoginForm
func validateLoginForm(form *LoginForm) (*ValidLoginForm, error) {

	vForm := &ValidLoginForm{Username: "", Email: "", Password: form.Password}
	// Validate username
	username, err := validateUsername(form.Username)
	if err != nil {
		// Validate email
		username, err = validateEmail(form.Username)
		if err != nil {
			return nil, ErrorLoginInfos
		}
		vForm.Email = username
	} else {
		vForm.Username = username
	}

	return vForm, nil
}

// validate bearer token
func validateBearerToken(tokenString string) (int64, string, error) {

	params := strings.Split(tokenString, " ")
	if params[0] != "Bearer" {
		return 0, "", ErrorTokenNoyValid
	}

	funcKey := func(*jwt.Token) (interface{}, error) {
		return TokenSignedString, nil
	}

	token, err := jwt.ParseWithClaims(params[1], &ClaimsJWT{}, funcKey)

	if err != nil {

	}

	if !token.Valid {
		return 0, "", ErrorTokenNoyValid

	}

	claims, ok := token.Claims.(*ClaimsJWT)

	if !ok {
		return 0, "", ErrorTokenNoyValid
	}

	return claims.ID, claims.Username, nil
}
