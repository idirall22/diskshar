package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) createUser(ctx context.Context, username, email, password string) error {

	email, err := validateEmail(email)
	if err != nil {
		return err
	}

	username, err = validateUsername(username)
	if err != nil {
		return err
	}

	passwordHased, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.provider.New(ctx, username, email, "", "", string(passwordHased), "")

	if err != nil {
		return err
	}
	return nil
}
