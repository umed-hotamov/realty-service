package providers

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/repository"

	"go.uber.org/zap"
)

type PropertyService struct {
	repo   repository.IPropertyRepo
	logger *zap.Logger
}

func NewPropertyService(repo repository.IPropertyRepo, logger *zap.Logger) *PropertyService {
	return &PropertyService{
		repo: repo,
    logger: logger,
	}
}

func (p *PropertyService) GetAll(ctx context.Context) ([]domain.Property, error) {
	properties, err := p.repo.GetAll(ctx)
	if err != nil {
		p.logger.Error("failed to get all properties", zap.Error(err))
		return nil, err
	}

	return properties, nil
}

func (p *PropertyService) GetUserProperties(ctx context.Context, userID domain.ID) ([]domain.Property, error) {
	properties, err := p.repo.GetUserProperties(ctx, userID)
	if err != nil {
		p.logger.Error("failed to get user properties", zap.Error(err))
		return nil, err
	}

	return properties, nil
}

func (p *PropertyService) Create(ctx context.Context, param domain.CreatePropertyParam) (domain.Property, error) {
	propertyDTO := domain.Property{}

	property, err := p.repo.Create(ctx, propertyDTO)
	if err != nil {
		p.logger.Error("failed to create property", zap.Error(err))
		return domain.Property{}, err
	}

	return property, nil
}

func (p *PropertyService) Update(ctx context.Context, param domain.UpdatePropertyParam, propertyID domain.ID) (domain.Property, error) {
	property, err := p.repo.GetPropertyByID(ctx, propertyID)
	if err != nil {
		p.logger.Error("failed to get property by id", zap.Error(err))
		return domain.Property{}, nil
	}

	property, err = p.repo.Update(ctx, property)
	if err != nil {
		p.logger.Error("failed to update property", zap.Error(err))
		return domain.Property{}, err
	}

	return property, nil
}

func (p *PropertyService) Delete(ctx context.Context, propertyID domain.ID) error {
	err := p.repo.Delete(ctx, propertyID)
	if err != nil {
		p.logger.Error("failed to delete property", zap.Error(err))
		return err
	}

	return nil
}
