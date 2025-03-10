package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/app/config"
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
  
  return handler
}
