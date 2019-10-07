package user

import (
	"context"
	"testing"
)

func testNew(t *testing.T) {
	// fmt.Println(testService.provider.(*))
	username := "Jhon"
	email := "Jhon@email.com"
	password := "password"

	user, err := testService.provider.New(context.Background(),
		username, email, "", "", password, "")

	if err != nil {
		t.Error("Error create new user, ", err)
		return
	}

	if user.Username != username {
		t.Error("Error username does not match")
	}
}
