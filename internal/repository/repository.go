package repository

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type IUserRepo interface {
  GetAll(ctx context.Context) ([]domain.User, error)
  GetByID(ctx context.Context) (domain.User, error)
  Create(ctx context.Context, user domain.User) (domain.User, error)
  Update(ctx context.Context, user domain.User) (domain.User, error)
  Delete(ctx context.Context, userID int) error
}

type IPropertyRepo interface {
  GetAll(ctx context.Context) ([]domain.Property, error)
  GetPropertyByAddress(ctx context.Context, address string) ([]domain.Property, error)
  GetUserProperty(ctx context.Context, userID int) ([]domain.Property, error)
  GetApprovedProperty(ctx context.Context) ([]domain.Property, error)
  GetFilteredProperty(ctx context.Context, filter domain.PropertyFilter) ([]domain.Property, error)
  Create(ctx context.Context, property domain.Property) (domain.Property, error)
  Update(ctx context.Context, property domain.Property) (domain.Property, error)
  Delete(ctx context.Context, property domain.Property) error
}

type IListingRepo interface {
  GetAll(ctx context.Context) ([]domain.Listing, error)
  GetUserListings(ctx context.Context, userID int) ([]domain.Listing, error)
  Create(ctx context.Context, listing domain.Listing) (domain.Listing, error)
  Update(ctx context.Context, listing domain.Listing) (domain.Listing, error)
  Delete(ctx context.Context, listingID int) error
}

type IApartmentBuildingRepo interface {
  GetAll(ctx context.Context) ([]domain.ApartmentBuilding, error)
  GatBuildingFlats(ctx context.Context, buildingID int) ([]domain.Flat, error)
  Create(ctx context.Context, building domain.ApartmentBuilding) (domain.ApartmentBuilding, error)
}
