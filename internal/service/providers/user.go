package providers

import (
	"context"

	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/repository"
	
  "go.uber.org/zap"
)

type UserService struct {
  repo   repository.IUserRepo
  logger *zap.Logger
}

func NewUserService(repo repository.IUserRepo) *UserService {
  return &UserService{
    repo: repo,
  }
}

func (u *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
  users, err := u.repo.GetAll(ctx)
  if err != nil {
    u.logger.Error("failed to get users", zap.Error(err))
    return nil, err
  }

  return users, nil
}

func (u *UserService) GetByID(ctx context.Context, id domain.ID) (domain.User, error) {
  user, err := u.repo.GetByID(ctx, id)
  if err != nil {
    u.logger.Error("failed to get user by ID", zap.Error(err))
    return domain.User{}, err
  }

  return user, nil
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (domain.User, error) {
  user, err := u.repo.GetByEmail(ctx, email)
  if err != nil {
    u.logger.Error("failed to get user by email", zap.Error(err))
    return domain.User{}, err
  }

  return user, nil
}

func (u *UserService) GetByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error) {
  users, err := u.repo.GetByRole(ctx, role)
  if err != nil {
    u.logger.Error("failed to get users", zap.Error(err))
    return nil, err
  }

  return users, nil
}

func (u *UserService) Create(ctx context.Context, param domain.CreateUserParam) (domain.User, error) {
  // need to add field check

  userDTO := domain.User{
    ID:       domain.NewID(),
    Role:     param.Role,
    Name:     param.Name,
    Surname:  param.Surname,
    Email:    param.Email,
    Password: param.Password,
    Phone:    param.Phone,
  }

  user, err := u.repo.Create(ctx, userDTO)
  if err != nil {
    u.logger.Error("failed to create user", zap.Error(err))
    return domain.User{}, err
  }

  return user, nil
}

func (u *UserService) Update(ctx context.Context, param domain.UpdateUserParam, userID domain.ID) (domain.User, error) {
  user, err := u.GetByID(ctx, userID)
  if err != nil {
    return domain.User{}, nil
  }

  user.Name = param.Name
  user.Surname = param.Surname
  user.Phone = param.Phone

  user, err = u.repo.Update(ctx, user)
  if err != nil {
    u.logger.Error("failed to update user", zap.Error(err))
    return domain.User{}, err
  }

  return user, nil
}

func (u *UserService) Delete(ctx context.Context, userID domain.ID) error {
  err := u.repo.Delete(ctx, userID)
  if err != nil {
    u.logger.Error("failed to delete user", zap.Error(err))
    return err
  }

  return nil
}
