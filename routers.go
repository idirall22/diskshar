package userAccount

import (
	"github.com/gorilla/mux"
)

// Router comment endpoints
func (s *Service) Router(r *mux.Router) {
	sr := r.PathPrefix("/accounts").Subrouter()
	sr.HandleFunc("/register", s.Register).Methods("POST")
	sr.HandleFunc("/login", s.Login).Methods("POST")
}
