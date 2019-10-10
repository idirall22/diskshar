package user

import (
	"context"
	"fmt"
	"testing"
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

func testAuthenticate(t *testing.T) {
	token, err := testService.Authenticate(context.Background(), username, "", userPassword)
	fmt.Println(err)
	fmt.Println(token)
}
