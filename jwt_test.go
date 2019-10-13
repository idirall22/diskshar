package userAccount

import (
	"testing"
)

// Test generate token
func TestGenerateToken(t *testing.T) {
	_, err := generateToken(1, "jhon")

	if err != nil {
		t.Error("Error should be nil but got: ", err)
	}
}
