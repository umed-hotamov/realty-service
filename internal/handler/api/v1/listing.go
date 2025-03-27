package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitListingRoutes(api *echo.Group) {
  listing := api.Group("/listings", h.verifyToken)
  {
    listing.GET("/", h.getAllListings)
    listing.GET("/:user_id", h.getUserListings)
    listing.POST("/create", h.createListing)
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

  return h.successResponse(c, listingsDTO)
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

 return h.successResponse(c, listingsDTO)
}

func (h *Handler) createListing(c echo.Context) error {
  userID, err := getIdFromRequestContext(c)
  if err != nil {
    return h.errorResponse(c, err)
  }

  var createListingDTO dto.CreateListingDTO
  err = c.Bind(&createListingDTO)
  if err != nil {
    return h.errorResponse(c, err)
  }

  _, err = h.listingService.Create(c.Request().Context(), domain.CreateListingParam{
    UserID:      userID,
    PropertyID:  domain.ID(createListingDTO.PropertyID),
    Title:       createListingDTO.Title,
    City:        createListingDTO.City,
    Offer:       createListingDTO.Offer,
    Price:       createListingDTO.Price,
    Description: createListingDTO.Description,
    Status:      createListingDTO.Status,
  })

  if err != nil {
    return h.errorResponse(c, err)
  }

  return h.successResponse(c, "created")
}

func (h *Handler) updateListing(c echo.Context) error {
  var updateListingDTO dto.UpdateListingDTO
  
  err := c.Bind(&updateListingDTO)
  if err != nil {
    return h.errorResponse(c, err)
  }

  _, err = h.listingService.Update(c.Request().Context(), domain.UpdateListingParam{
    Title:       updateListingDTO.Title,
    City:        updateListingDTO.City,
    Offer:       updateListingDTO.Offer,
    Price:       updateListingDTO.Price,
    Description: updateListingDTO.Description,
  }, domain.ID(updateListingDTO.ID))

  if err != nil {
    return h.errorResponse(c, err)
  }
  
  return h.successResponse(c, "updated")
}
