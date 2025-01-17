package repository

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
)

type IFlatRepo interface {
  GetFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error)
  GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error)
  Create(ctx context.Context, flat domain.Flat) (domain.Flat, error)
  Update(ctx context.Context, flatID int) (domain.Flat, error)
}
