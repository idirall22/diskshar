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
