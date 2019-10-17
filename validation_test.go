package userAccount

import (
	"testing"

	"github.com/idirall22/user/models"
)

// Test email
func TestValidateEmail(t *testing.T) {
	emailsTest := []string{
		"jhon@email.com",
		"@email.com",
	}

	for i := 0; i < len(emailsTest); i++ {
		_, err := validateEmail(emailsTest[i])

		switch i {
		case 0:
			if err != nil {
				t.Error("Error this error should be nil but got :", err)
			}
		case 1:
			if err == nil {
				t.Error("Error this error should not be nil")
			}
		}
	}

}

// Test email
func TestValidateUsername(t *testing.T) {
	usernameTest := []string{
		"Jhon",
		"@jhon+",
	}

	for i := 0; i < len(usernameTest); i++ {
		_, err := validateUsername(usernameTest[i])

		switch i {
		case 0:
			if err != nil {
				t.Error("Error this error should be nil but got :", err)
			}
		case 1:
			if err == nil {
				t.Error("Error this error should not be nil")
			}
		}
	}

}

// validate register form
func TestValidateRegisterForm(t *testing.T) {
	registerForm := []*models.RegisterForm{
		// When form is valid
		{Username: "jhon", Email: "jhon@gmail.com", Password: testPassword},
		// When form is not valid "username"
		{Username: "@jhon", Email: "jhon@gmail.com", Password: testPassword},
		// When form is not valid "email"
		{Username: "jhon", Email: "jhongmail.com", Password: testPassword},
	}

	for i, form := range registerForm {
		_, err := validateRegisterForm(form)
		switch i {
		case 0:
			if err != nil {
				t.Error("Error this error should be nil but got :", err)

			}
		case 1:
			if err != ErrorUsernameNotValid {
				t.Error("Error this error should be ErrorUsernameNotValid but got :", err)
			}
		case 2:
			if err != ErrorEmailNotValid {
				t.Error("Error this error should be ErrorEmailNotValid but got :", err)
			}
		}
	}
}

// Test validate login form
func TestValidateLoginForm(t *testing.T) {
	loginForm := []*models.LoginForm{
		// When form user use username to login
		{Username: "jhon", Password: "password"},

		// When form user use username to email
		{Username: "jhon@email.com", Password: "password"},

		// When form user use bad credentials
		{Username: "jhonemail.com", Password: "password"},
	}
	for i, form := range loginForm {
		_, err := validateLoginForm(form)

		switch i {
		case 0:
			if err != nil {
				t.Error("Error this error should be nil but got: ", err)
			}
		case 1:
			if err != nil {
				t.Error("Error this error should be nil but got: ", err)
			}
		case 2:
			if err != ErrorLoginInfos {
				t.Errorf("Error this error should be %s but got: %s\n", ErrorLoginInfos.Error(), err)
			}
		}
	}
}

// Test validate token
func testValidateToken(t *testing.T) {
	testTokens := []string{
		testTokenString,
		testTokenString[:len(testTokenString)-5],
	}

	for i, token := range testTokens {
		_, _, err := validateBearerToken(token)
		switch i {
		case 0:
			if err != nil {
				t.Error("Error should be nil but got: ", err)
				break
			}
		case 1:
			if err == nil {
				t.Errorf("Error should be \"%s\" but got: \"%s\"", ErrorTokenNoyValid, err)
				break
			}
		}
	}

}
