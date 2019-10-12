package user

import (
	"fmt"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Minimum length of password
	minPassLength = 8

	// Maximum length of password
	maxPassLength = 128

	// Default entropy to compare with password
	defaultEntropy = float64(40)
)

var (
	// ErrorMinPass when password length is less then minPassLength
	ErrorMinPass = fmt.Errorf("Minimum password length must be %d", minPassLength)

	// ErrorMaxPass when password length is gt maxPassLength
	ErrorMaxPass = fmt.Errorf("Maximum password length must be %d", maxPassLength)

	// ErrorPassStrength when password strength is less then defaultEntropy
	ErrorPassStrength = fmt.Errorf("Maximum password length must be %d", maxPassLength)
)

// validate password
func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// validate password
func validatePassword(password string) error {

	if len(password) < minPassLength {
		return ErrorMinPass
	}

	if len(password) > maxPassLength {
		return ErrorMaxPass
	}

	// check password strength
	entropy := zxcvbn.PasswordStrength(password, nil).Entropy

	if entropy < defaultEntropy {
		return ErrorPassStrength
	}
	return nil
}
