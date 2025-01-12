package repository

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
)

type PostgresHouseRepo struct {
}

func NewHouseRepo() *PostgresHouseRepo {
  return &PostgresHouseRepo{
  }
}

func (p *PostgresHouseRepo) GetAll(ctx context.Context) ([]domain.House, error) {
  var houses []domain.House
  return houses, nil
}

func (p *PostgresHouseRepo) Create(ctx context.Context, house domain.House) (domain.House, error) {
  return domain.House{}, nil
}
