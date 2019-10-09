package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// make request for tests
func makeTestRequest(t *testing.T, f http.HandlerFunc, method, url string, data []byte) *httptest.ResponseRecorder {

	body := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r, err := http.NewRequest(method, url, body)
	if err != nil {

	}
	handler := http.HandlerFunc(f)
	handler.ServeHTTP(w, r)
	return w
}

// test register handler
func testRegister(t *testing.T) {

	registerForm := []RegisterForm{
		{Username: "alice1", Email: "alice@gmail.com", Password: "password"},
		{Username: "alice1", Email: "alice3@gmail.com", Password: "password"},
		{Username: "alice2", Email: "alice@gmail.com", Password: "password"},
	}

	for i, form := range registerForm {
		data, err := json.Marshal(form)

		if err != nil {
			t.Error(err)
			return
		}
		resp := makeTestRequest(t, testService.Register, "POST", "/register", data)

		switch i {
		case 0:
			if resp.Code != http.StatusOK {
				t.Errorf("Error status should be %d but got %d", resp.Code, http.StatusOK)
			}
		case 1:
			if resp.Code != http.StatusConflict {
				t.Errorf("Error status should be %d but got %d", resp.Code, http.StatusConflict)
			}
		case 2:
			if resp.Code != http.StatusConflict {
				t.Errorf("Error status should be %d but got %d", resp.Code, http.StatusConflict)
			}
		}
	}
}
