package repository

import (
	"context"
	"log"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type PostgresHouseRepo struct {
}

func NewHouseRepo() *PostgresHouseRepo {
  return &PostgresHouseRepo{
  }
}

func (p *PostgresHouseRepo) GetAll(ctx context.Context) ([]domain.House, error) {
  var houses []domain.House
  log.Print("Got all houses")
  return houses, nil
}

func (p *PostgresHouseRepo) Create(ctx context.Context, house domain.House) (domain.House, error) {
  log.Print("House created")
  return domain.House{}, nil
}
