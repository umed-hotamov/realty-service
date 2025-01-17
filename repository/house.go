package repository

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
)

type IHouseRepo interface {
  GetAll(ctx context.Context) ([]domain.House, error)
  Create(ctx context.Context, house domain.House) (domain.House, error)
}
