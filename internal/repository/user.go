package repository

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
)

type IUserRepo interface {
  Create(ctx context.Context, user domain.User) (domain.User, error)
  Update(ctx context.Context, user domain.User) (domain.User, error)
}
