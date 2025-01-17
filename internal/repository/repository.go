package repository

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type IFlatRepo interface {
  GetFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error)
  GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error)
  Create(ctx context.Context, flat domain.Flat) (domain.Flat, error)
  Update(ctx context.Context, flatID int) (domain.Flat, error)
}

type IHouseRepo interface {
  GetAll(ctx context.Context) ([]domain.House, error)
  Create(ctx context.Context, house domain.House) (domain.House, error)
}
