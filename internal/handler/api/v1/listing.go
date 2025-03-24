package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitListingRoutes(api *echo.Group) {
  listings := api.Group("/listings")
  {
    listings.GET("/", h.getAllListings)
    listings.GET("/:user_id", h.getUserListings)
    
  }
}

func (h *Handler) getAllListings(c echo.Context) error {
  listings, err := h.listingService.GetAll(c.Request().Context())
  if err != nil {
    return h.errorResponse(c, err)
  }

  listingsDTO := make([]dto.ListingDTO, len(listings))
  for i, listing := range listings {
    listingsDTO[i] = dto.NewListingDTO(listing)
  }

  return h.successResponse(c, listings)
}
  
func (h *Handler) getUserListings(c echo.Context) error {
  userID, err := getIDFromPath(c, "user_id")
  if err != nil {
    return h.errorResponse(c, err)
  }

  listings, err := h.listingService.GetUserListings(c.Request().Context(), userID)
  if err != nil {
    return h.errorResponse(c, err)
  }

  listingsDTO := make([]dto.ListingDTO, len(listings))
  for i, listing := range listings {
    listingsDTO[i] = dto.NewListingDTO(listing)
  }

  return h.successResponse(c, listings)
}

// func (h *Handler) postListing(c echo.Context) error {
// }
