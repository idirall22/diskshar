package user

import (
	"encoding/json"
	"net/http"
)

// Register endpoint: "auth/register"
func (s *Service) Register(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	form := &RegisterForm{}
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	validatedForm, err := validateLoginForm(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	err = s.createUser(r.Context(),
		validatedForm.Username,
		validatedForm.Email,
		validatedForm.Password,
	)

	if err != nil {
		switch err {
		case ErrorEmailTaken:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		case ErrorUsernameTaken:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, "Error server", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
}

// Login endpoint: "auth/login"
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {

}
