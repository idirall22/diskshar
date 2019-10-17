package userAccount

import (
	"context"
	"testing"

	"github.com/idirall22/user/models"
)

// Test create user
func testCreateUser(t *testing.T) {

	testUsers := []struct {
		username string
		email    string
	}{
		{"jane", "jane@email.com"},
		{"jane", "jane2@email.com"},
		{"jane2", "jane@email.com"},
	}

	for i, user := range testUsers {
		err := testService.createUser(context.Background(), user.username, user.email, userPassword)

		switch i {
		case 0:
			if err != nil {
				t.Error("Error: the error should be nil and got :", err)
			}
		case 1:
			if err != ErrorUsernameTaken {
				t.Errorf("Error: the error should be \"%s\" and got :%s\n", ErrorUsernameTaken, err)
			}
		case 2:
			if err != ErrorEmailTaken {
				t.Errorf("Error: the error should be \"%s\" and got :%s\n", ErrorEmailTaken, err)
			}
		}
	}
}

// Test Authenticate user
func testAuthenticate(t *testing.T) {

	testUsers := []models.ValidLoginForm{
		// When user use username to login
		{Username: "jane", Email: "", Password: userPassword},
		// When user use email to login
		{Username: "", Email: "jane@email.com", Password: userPassword},
		// When user use username but this one does not exists
		{Username: "Xman", Email: "", Password: userPassword},
	}

	for i, user := range testUsers {
		_, err := testService.Authenticate(context.Background(), user.Username, user.Email, user.Password)

		switch i {
		case 0:
			if err != nil {
				t.Error("Error should be nil but got: ", err)
			}
		case 1:
			if err != nil {
				t.Error("Error should be nil but got: ", err)
			}
		case 2:
			if err != ErrorNoUser {
				t.Errorf("Error should be \"%s\" but got: \"%s\"\n", ErrorNoUser, err)
			}
		}
	}
}
