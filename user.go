package user

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrorUsernameTaken when username already taken
	ErrorUsernameTaken = errors.New("Username already taken")
	// ErrorEmailTaken when email already taken
	ErrorEmailTaken = errors.New("Email already taken")
)

// create user
func (s *Service) createUser(ctx context.Context, username, email, password string) error {

	passwordHased, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.provider.New(ctx, username, email, "", "", string(passwordHased), "")

	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			if e.Code == "23505" {
				return parseUniqueConstraintError(e.Constraint)
			}
		}
		return err
	}
	return nil
}
