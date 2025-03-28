package manager

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
	"github.com/umed-hotamov/realty-service/internal/app/config"
	"github.com/umed-hotamov/realty-service/internal/authHelpers/storage"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/errs"
)

type Config struct {
	Secret           string
	AccessTokenTime  int64
	RefreshTokenTime int64
}

type AuthManager struct {
	cfg            *config.Config
	sessionStorage storage.SessionStorage
}

func NewAuthManager(cfg *config.Config, sessionStorage storage.SessionStorage) *AuthManager {
	return &AuthManager{
		cfg:            cfg,
		sessionStorage: sessionStorage,
	}
}

func (p *AuthManager) CreateJWTSession(payload domain.AuthPayload,
	fingerprint string) (domain.AuthDetails, error) {
	accessExpTime := time.Minute * time.Duration(p.cfg.Jwt.AccessTokenTime)
	accessExp := time.Now().Add(accessExpTime).Unix()
	claims := jwt.MapClaims{
		"exp":    accessExp,
		"userID": payload.UserID.String(),
		"role":   payload.Role,
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := unsignedToken.SignedString([]byte(p.cfg.Jwt.Secret))
	if err != nil {
		return domain.AuthDetails{}, err
	}

	refreshToken := uuid.NewV4().String()
	refreshExpTime := time.Minute * time.Duration(p.cfg.Jwt.RefreshTokenTime)
	refreshExp := time.Now().Add(refreshExpTime).Unix()

	session := storage.AuthSession{
		RefreshToken: refreshToken,
		RefreshExp:   refreshExp,
		Fingerprint:  fingerprint,
		Payload:      payload,
	}

	err = p.sessionStorage.Put(refreshToken, session, refreshExpTime)
	if err != nil {
		return domain.AuthDetails{}, err
	}

	return domain.AuthDetails{
		AccessToken:  domain.Token(accessToken),
		RefreshToken: domain.Token(refreshToken),
	}, nil
}

func (p *AuthManager) RefreshJWTSession(refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	session, err := p.sessionStorage.Get(refreshToken.String())
	if err != nil {
		return domain.AuthDetails{}, errors.Wrap(errs.ErrInvalidTokenClaims, err.Error())
	}

	err = p.sessionStorage.Delete(refreshToken.String())
	if err != nil {
		return domain.AuthDetails{}, err
	}

	if session.Fingerprint != fingerprint {
		return domain.AuthDetails{}, errs.ErrInvalidFingerprint
	}

	return p.CreateJWTSession(session.Payload, fingerprint)
}

func (p *AuthManager) DeleteJWTSession(refreshToken domain.Token) error {
	return p.sessionStorage.Delete(refreshToken.String())
}

func (p *AuthManager) VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error) {
	token, err := jwt.Parse(accessToken.String(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return domain.AuthPayload{}, errors.Wrap(errs.ErrInvalidToken, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		payload := domain.AuthPayload{
			UserID: domain.ID(claims["userID"].(string)),
			Role:   domain.UserRole(claims["role"].(float64)),
		}
		return payload, nil
	}

	return domain.AuthPayload{}, errs.ErrInvalidTokenClaims
}

func (p *AuthManager) GenPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(p.cfg.Jwt.Secret)))
}
