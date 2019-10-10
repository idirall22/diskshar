package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrorUsernameTaken when username already taken
	ErrorUsernameTaken = errors.New("Username already taken")

	// ErrorEmailTaken when email already taken
	ErrorEmailTaken = errors.New("Email already taken")

	// ErrorPassword when password is not valid
	ErrorPassword = errors.New("Password not valid")

	// ErrorNoUser when there is no user with credentials used
	ErrorNoUser = errors.New("There is no user with credentials used")
)

// create user
func (s *Service) createUser(ctx context.Context, username, email, password string) error {

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = s.provider.New(ctx, username, email, "", "", string(passwordHashed), "")

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

// Authenticate auth a user
func (s *Service) Authenticate(ctx context.Context, username, email, password string) (string, error) {

	user, err := s.provider.Get(ctx, 0, username, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return "", ErrorNoUser
		}
		return "", err
	}
	err = validatePassword(user.Password, password)

	if err != nil {
		return "", ErrorPassword
	}

	return generateToken(user.ID, user.Username)
}
