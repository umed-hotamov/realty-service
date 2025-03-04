package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) errorResponse(c echo.Context, err error) error {
  return c.JSON(http.StatusInternalServerError, err)
}

func (h *Handler) successResponse(c echo.Context, data interface{}) error {
  return c.JSON(http.StatusOK, data)
}
