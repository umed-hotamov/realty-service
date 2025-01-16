package repository

import (
	"context"
	"log"

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
  log.Print("Got flats")
  return flats, nil
}

func (p *PostgresFlatRepo) GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
  var approvedFlats []domain.Flat
  log.Print("Got approved flats")
  return approvedFlats, nil
}

func (p *PostgresFlatRepo) Create(ctx context.Context, flat domain.Flat) (domain.Flat, error) {
  log.Print("Flat created")
  return domain.Flat{}, nil
}

func (p *PostgresFlatRepo) Update(ctx context.Context, flatID int) (domain.Flat, error) {
  log.Printf("Flat updated")
  return domain.Flat{}, nil
}
