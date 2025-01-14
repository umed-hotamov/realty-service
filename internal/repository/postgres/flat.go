package repository

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
)

type PostgresFlatRepo struct {
}

func NewFlatRepo() *PostgresFlatRepo {
  return &PostgresFlatRepo{
  }
}

func (p *PostgresFlatRepo) GetFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
  var flats []domain.Flat
  return flats, nil
}

func (p *PostgresFlatRepo) GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
  var approvedFlats []domain.Flat
  return approvedFlats, nil
}

func (p *PostgresFlatRepo) Create(ctx context.Context, flat domain.Flat) (domain.Flat, error) {
  return domain.Flat{}, nil
}
