package providers

import (
	"context"

	manager "github.com/umed-hotamov/realty-service/internal/authHelpers/authManager"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/errs"
	"github.com/umed-hotamov/realty-service/internal/service"
	"go.uber.org/zap"
)

type AuthService struct {
  authManager *manager.AuthManager 
  userService service.IUserService 
  logger *zap.Logger
}

func NewAuthService(authManager *manager.AuthManager, userService service.IUserService) *AuthService {
	return &AuthService{
		authManager: authManager,
		userService:  userService,
	}
}

func (a *AuthService) SignIn(ctx context.Context, param domain.SignInParam) (domain.AuthDetails, error) {
	user, err := a.userService.GetByEmail(ctx, param.Email)
	if err != nil {
		return domain.AuthDetails{}, errs.ErrEmail
	}

	if user.Password != a.authManager.GenPasswordHash(param.Password) {
		return domain.AuthDetails{}, errs.ErrPassword
	}
	ad, err := a.authManager.CreateJWTSession(domain.AuthPayload{UserID: user.ID, Role: user.Role}, param.Fingerprint)
	if err != nil {
		return domain.AuthDetails{}, err
	}

	return ad, nil
}

func (a *AuthService) SignUp(ctx context.Context, param domain.SignUpParam) (domain.User, error) {
	user, err := a.userService.Create(ctx, domain.CreateUserParam{
    Role:     param.Role,
    Name:     param.Name,
    Surname:  param.Surname,
    Email:    param.Email,
    Password: a.authManager.GenPasswordHash(param.Password),
    Phone:    param.Phone,
	})

	if err != nil {
    a.logger.Error("failed to signUp user", zap.Error(err))
		return domain.User{}, err
	}

	return user, nil
}

func (a *AuthService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	err := a.authManager.DeleteJWTSession(refreshToken)
	if err != nil {
    a.logger.Error("faild to logout: %s", zap.Error(err))
		return err
	}

	return nil
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	ad, err := a.authManager.RefreshJWTSession(refreshToken, fingerprint)
	if err != nil {
		return domain.AuthDetails{}, err
	}

	return ad, nil
}

func (a *AuthService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	ap, err := a.authManager.VerifyJWTToken(accessToken)
	if err != nil {
		return domain.AuthPayload{}, err
	}

	return ap, nil
}
