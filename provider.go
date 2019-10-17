package userAccount

import (
	"context"

	"github.com/idirall22/user/models"
)

// Provider interface
type Provider interface {
	New(context.Context, string, string, string, string, string, string) (*models.User, error)
	Get(context.Context, int64, string, string) (*models.User, error)
	Update(context.Context, int64, string, string, string) error
	Delete(ctx context.Context, id int64) error
}
