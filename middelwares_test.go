package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test AuthnticateUser
func testAuthnticateUser(t *testing.T) {

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authorized view"))

	})

	testTokens := []string{

		// when token is valid
		testTokenString,

		// when token is not valid
		testTokenString[:len(testTokenString)-5],

		// when there is not a token in header
		"",
	}

	for i, token := range testTokens {

		resp := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/", nil)
		r.Header.Add("Authorization", token)
		handler := AuthnticateUser(testHandler)

		handler.ServeHTTP(resp, r)

		switch i {
		case 0:
			if resp.Code != http.StatusOK {
				t.Error("Error should be nil but got:", resp.Code)
				break
			}
		case 1:
			if resp.Code != http.StatusBadRequest {
				t.Errorf("Error should be \"%d\" but got: \"%d\"\n", http.StatusBadRequest, resp.Code)
				break
			}
		case 2:
			if resp.Code != http.StatusUnauthorized {
				t.Errorf("Error should be \"%d\" but got: \"%d\"\n", http.StatusUnauthorized, resp.Code)
				break
			}
		}
	}
}

// Test storeUserInContext
func TestStoreUserInContext(t *testing.T) {
	ctx := storeUserInContext(context.Background(), 1, "username")

	if ctx.Value(IDCtx) != int64(1) {
		t.Error("Error context does not containe user id")
	}

	if ctx.Value(UsernameCtx) != "username" {
		t.Error("Error context does not containe user username")
	}
}
