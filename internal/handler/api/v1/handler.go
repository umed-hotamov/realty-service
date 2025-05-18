package v1

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/app/config"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/service"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Handler struct {
	config          *config.Config
	logger          *zap.Logger
	userService     service.IUserService
	propertyService service.IPropertyService
	listingService  service.IListingService
	authService     service.IAuthService
}

type HandlerParams struct {
	fx.In

	Config          *config.Config
	Logger          *zap.Logger
	UserService     service.IUserService
	PropertyService service.IPropertyService
	ListingService  service.IListingService
}

func NewHandler(params HandlerParams, e *echo.Echo) *Handler {
	handler := &Handler{
		config:          params.Config,
		logger:          params.Logger,
		userService:     params.UserService,
		propertyService: params.PropertyService,
		listingService:  params.ListingService,
	}

	api := e.Group("/api/v1")
	handler.InitUserRoutes(api)
	handler.InitPropertyRoutes(api)
	handler.InitListingRoutes(api)

	return handler
}

func getIDFromPath(c echo.Context, param string) (domain.ID, error) {
	idString := c.Param(param)
	if idString == "" {
		return "", ErrorEmptyParam
	}

	if _, err := uuid.Parse(idString); err != nil {
		return "", ErrorInvalidUUID
	}

	return domain.ID(idString), nil
}

func getIdFromRequestContext(c echo.Context) (domain.ID, error) {
	id := c.Get("userID")
	if id == nil {
		return "", ErrorUnauthorized
	}

	return domain.ID(id.(string)), nil
}

func extractAuthToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if headerParts[1] == "" {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

func (h *Handler) verifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := extractAuthToken(c)
		if err != nil {
			return h.errorResponse(c, ErrorUnauthorized)
		}

		payload, err := h.authService.Payload(c.Request().Context(), domain.Token(tokenString))
		if err != nil {
			return h.errorResponse(c, err)
		}

		c.Set("userID", payload.UserID.String())
		c.Set("role", int(payload.Role))

		return next(c)
	}
}
