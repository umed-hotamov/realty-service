package providers

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/repository"

	"go.uber.org/zap"
)

type ListingService struct {
	repo   repository.IListingRepo
	logger *zap.Logger
}

func NewListingService(repo repository.IListingRepo) *ListingService {
	return &ListingService{
		repo: repo,
	}
}

func (l *ListingService) GetAll(ctx context.Context) ([]domain.Listing, error) {
	listings, err := l.GetAll(ctx)
	if err != nil {
		l.logger.Error("failed to ger all listings", zap.Error(err))
		return nil, err
	}

	return listings, nil
}

func (l *ListingService) GetUserListings(ctx context.Context, userID domain.ID) ([]domain.Listing, error) {
	listings, err := l.GetUserListings(ctx, userID)
	if err != nil {
		l.logger.Error("failed to ger user listings", zap.Error(err))
		return nil, err
	}

	return listings, nil
}

func (l *ListingService) Create(ctx context.Context, param domain.CreateListingParam) (domain.Listing, error) {
	listingDTO := domain.Listing{
		ID:          domain.NewID(),
		UserID:      param.UserID,
		PropertyID:  param.PropertyID,
		Title:       param.Title,
		City:        param.City,
		Offer:       param.Offer,
		Price:       param.Price,
		Description: param.Description,
		Status:      param.Status,
	}

	listing, err := l.repo.Create(ctx, listingDTO)
	if err != nil {
		l.logger.Error("failed to create listing", zap.Error(err))
		return domain.Listing{}, err
	}

	return listing, nil
}

func (l *ListingService) Update(ctx context.Context, param domain.UpdateListingParam, listingID domain.ID) (domain.Listing, error) {
	listing, err := l.repo.GetListingByID(ctx, listingID)
	if err != nil {
		l.logger.Error("failed to get listing by id", zap.Error(err))
		return domain.Listing{}, nil
	}

	listing.Title = param.Title
	listing.City = param.City
	listing.Offer = param.Offer
	listing.Price = param.Price
	listing.Description = param.Description

	listing, err = l.repo.Update(ctx, listing)
	if err != nil {
		l.logger.Error("failed to update listing", zap.Error(err))
		return domain.Listing{}, err
	}

	return listing, nil
}

func (l *ListingService) Delete(ctx context.Context, listingID domain.ID) error {
	err := l.repo.Delete(ctx, listingID)
	if err != nil {
		l.logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	return nil
}
