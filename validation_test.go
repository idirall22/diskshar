package user

import (
	"testing"
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
func TestValidateLoginForm(t *testing.T) {
	registerForm := []*RegisterForm{
		// When form is valid
		{Username: "jhon", Email: "jhon@gmail.com", Password: "password"},
		// When form is not valid "username"
		{Username: "@jhon", Email: "jhon@gmail.com", Password: "password"},
		// When form is not valid "email"
		{Username: "jhon", Email: "jhongmail.com", Password: "password"},
	}

	for i, form := range registerForm {
		_, err := validateLoginForm(form)
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
