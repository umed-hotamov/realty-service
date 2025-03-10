package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitListingRoutes(api *echo.Group) {
  listings := api.Group("/listings")
  {
    listings.GET("/", h.getAllListings)
    listings.GET("/:listing_id", h.getListing)
    listings.POST("/post")
  }
}

func (h *Handler) getAllListings(c echo.Context) error {
}

func (h *Handler) getListing(c echo.Context) error {

}

func (h *Handler) postListing(c echo.Context) error {

}
