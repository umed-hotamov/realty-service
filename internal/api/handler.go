package api

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/service"

	"go.uber.org/fx"
)

type Handler struct {
  flatService  service.IFlatService
  houseService service.IHouseService
}

type HandlerParams struct {
  fx.In
  FlatService  service.IFlatService
  HouseService service.IHouseService
}

func NewHandler(params HandlerParams, router *echo.Echo) *Handler {
  handler := &Handler{
    flatService:  params.FlatService,
    houseService: params.HouseService,
  }



  return handler
}
