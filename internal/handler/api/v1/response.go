package v1

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	BadRequest          = "bad request"
	NotFound            = "not found"
	Unauthorized        = "unauthorized"
	Forbidden           = "forbidden"
	InternalServerError = "internal server error"
	RequestTimeout      = "request timeout"
)

var (
	ErrorBadRequest   = errors.New("bad request")
	ErrorUnauthorized = errors.New("unauthorized")
	ErrorForbidden    = errors.New("forbidden")
	ErrorEmptyParam   = errors.New("empty query parameter")
	ErrorInvalidUUID  = errors.New("invalid uuid")
)

func (h *Handler) errorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, err)
}

func (h *Handler) successResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
