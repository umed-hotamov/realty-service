package repository

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type IUserRepo interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id domain.ID) (domain.User, error)
	GetByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}

type IPropertyRepo interface {
	GetAll(ctx context.Context) ([]domain.Property, error)
	GetPropertyByID(ctx context.Context, propertyID domain.ID) (domain.Property, error)
	GetUserProperties(ctx context.Context, userID domain.ID) ([]domain.Property, error)
	Create(ctx context.Context, property domain.Property) (domain.Property, error)
	Update(ctx context.Context, property domain.Property) (domain.Property, error)
	Delete(ctx context.Context, propertyID domain.ID) error
}

type IListingRepo interface {
	GetAll(ctx context.Context) ([]domain.Listing, error)
	GetListingByID(ctx context.Context, userID domain.ID) (domain.Listing, error)
	GetUserListings(ctx context.Context, userID domain.ID) ([]domain.Listing, error)
	Create(ctx context.Context, listing domain.Listing) (domain.Listing, error)
	Update(ctx context.Context, listing domain.Listing) (domain.Listing, error)
	Delete(ctx context.Context, listingID domain.ID) error
}
