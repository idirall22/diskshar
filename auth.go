package user

import (
	"context"
)

// Authentication interface
type Authentication interface {
	login(context.Context, string, string, string) (string, error)
	register(context.Context, string, string, string) error
}
