package service

import (
	"context"

	"github.com/umed-hotamov/house-rental/internal/domain"
	"github.com/umed-hotamov/house-rental/internal/repository"
)

type FlatService struct {
  repo repository.IFlatRepo
}

func NewFlatService(repo repository.IFlatRepo) *FlatService {
  return &FlatService{
    repo: repo,
  }
}

func (f *FlatService) GetFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
  flats, err := f.repo.GetFlatsByHouseID(ctx, houseID)
  if err != nil {
    return nil, err
  }

  return flats, nil
} 

func (f *FlatService) GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
  flats, err := f.repo.GetApprovedFlatsByHouseID(ctx, houseID)
  if err != nil {
    return nil, err
  }

  return flats, nil
}

func (f *FlatService) Create(ctx context.Context, flat domain.Flat) (domain.Flat, error) {
  flat, err := f.repo.Create(ctx, flat)
  if err != nil {
    return domain.Flat{}, err
  }

  return flat, nil
}

func (f *FlatService) Update(ctx context.Context, flatID int) (domain.Flat, error) {
  newFlat, err := f.repo.Update(ctx, flatID)
  if err != nil {
    return domain.Flat{}, err
  }

  return newFlat, nil
}
