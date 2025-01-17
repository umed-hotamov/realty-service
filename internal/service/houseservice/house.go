package service

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/repository"
)

type HouseService struct {
  repo repository.IHouseRepo
}

func NewHouseService(repo repository.IHouseRepo) *HouseService {
  return &HouseService{
    repo: repo,
  }
}

func (hs *HouseService) GetAll(ctx context.Context) ([]domain.House, error) {
  houses, err := hs.repo.GetAll(ctx)
  if err != nil {
    return nil, err
  }

  return houses, nil
}

func (hs *HouseService) Create(ctx context.Context, house domain.House) (domain.House, error) {
  house, err := hs.repo.Create(ctx, house)
  if err != nil {
    return domain.House{}, err
  }

  return house, err
}
