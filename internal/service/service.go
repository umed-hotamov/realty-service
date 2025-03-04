package service

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type IUserService interface {
  GetAll(ctx context.Context) ([]domain.User, error)
  GetByID(ctx context.Context, id domain.ID) (domain.User, error)
  GetByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error)
  Create(ctx context.Context, param domain.CreateUserParam) (domain.User, error)
  Update(ctx context.Context, param domain.UpdateUserParam, userID domain.ID) (domain.User, error)
  Delete(ctx context.Context, userID domain.ID) error
}

type IPropertyService interface {
  GetAll(ctx context.Context) ([]domain.Property, error)
  GetPropertiesByAddress(ctx context.Context, address string) ([]domain.Property, error)
  GetUserProperties(ctx context.Context, userID domain.ID) ([]domain.Property, error)
  GetPropertyDetails(ctx context.Context, propertyID domain.ID) (domain.Property, error)
  Create(ctx context.Context, property domain.Property) (domain.Property, error)
  Update(ctx context.Context, property domain.Property) (domain.Property, error)
  Delete(ctx context.Context, property domain.Property) error
}

type IListingService interface {
  GetAll(ctx context.Context) ([]domain.Listing, error)
  GetUserListings(ctx context.Context, userID int) ([]domain.Listing, error)
  Create(ctx context.Context, listing domain.Listing) (domain.Listing, error)
  Update(ctx context.Context, listing domain.Listing) (domain.Listing, error)
  Delete(ctx context.Context, listingID int) error
}
