package service

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
)

type IUserService interface {
  GetAll(ctx context.Context) ([]domain.User, error)
  GetByID(ctx context.Context, id domain.ID) (domain.User, error)
  GetByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error)
  GetByEmail(ctx context.Context, email string) (domain.User, error)
  Create(ctx context.Context, param domain.CreateUserParam) (domain.User, error)
  Update(ctx context.Context, param domain.UpdateUserParam, userID domain.ID) (domain.User, error)
  Delete(ctx context.Context, userID domain.ID) error
}

type IPropertyService interface {
  GetAll(ctx context.Context) ([]domain.Property, error)
  GetUserProperties(ctx context.Context, userID domain.ID) ([]domain.Property, error)
  Create(ctx context.Context, param domain.CreatePropertyParam) (domain.Property, error)
  Update(ctx context.Context, param domain.UpdatePropertyParam, propertyID domain.ID) (domain.Property, error)
  Delete(ctx context.Context, propertyID domain.ID) error
}

type IListingService interface {
  GetAll(ctx context.Context) ([]domain.Listing, error)
  GetUserListings(ctx context.Context, userID domain.ID) ([]domain.Listing, error)
  Create(ctx context.Context, param domain.CreateListingParam) (domain.Listing, error)
  Update(ctx context.Context, param domain.UpdateListingParam, listingID domain.ID) (domain.Listing, error)
  Delete(ctx context.Context, listingID domain.ID) error
}

type IAuthService interface {
	SignIn(ctx context.Context, param domain.SignInParam) (domain.AuthDetails, error)
	SignUp(ctx context.Context, param domain.SignUpParam) error
	LogOut(ctx context.Context, refreshToken domain.Token) error
	Refresh(ctx context.Context, refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)
	Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error)
}
