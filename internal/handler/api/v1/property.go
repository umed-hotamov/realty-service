package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitPropertyRoutes(api *echo.Group) {
	property := api.Group("/property", h.verifyToken)
	{
		property.GET("/:user_id", h.getUserProperties)
	  property.POST("", h.createProperty)
  }
}

func (h *Handler) getUserProperties(c echo.Context) error {
	userID, err := getIDFromPath(c, "user_id")
	if err != nil {
		return h.errorResponse(c, err)
	}
  
	properties, err := h.propertyService.GetUserProperties(c.Request().Context(), userID)
	if err != nil {
		return h.errorResponse(c, err)
	}

	propertiesDTO := make([]dto.PropertyDTO, len(properties))
	for i, property := range properties {
		propertiesDTO[i] = dto.NewPropertyDTO(property)
	}

	return h.successResponse(c, propertiesDTO)
}

func (h *Handler) createProperty(c echo.Context) error {
	userID, err := getIdFromRequestContext(c)
	if err != nil {
		return h.errorResponse(c, err)
	}

	var createPropertyDTO dto.CreatePropertyDTO
	err = c.Bind(&createPropertyDTO)
	if err != nil {
		return h.errorResponse(c, err)
	}

	var buildingID domain.ID
	if createPropertyDTO.BuildingID != nil && *createPropertyDTO.BuildingID != "" {
		buildingID = domain.ID(*createPropertyDTO.BuildingID)
	} else {
		buildingID = ""
	}

	var houseType *domain.HouseType
	if createPropertyDTO.HouseType != nil {
		houseType = createPropertyDTO.HouseType
	} else {
		houseType = nil
	}
	_, err = h.propertyService.Create(c.Request().Context(), domain.CreatePropertyParam{
		OwnerID:      userID,
		BuildingID:   buildingID,
		PropertyType: createPropertyDTO.PropertyType,
		HouseType:    *houseType,
		PropertyDetails: domain.PropertyDetails{
			PropertyID:      domain.ID(createPropertyDTO.PropertyDetails.PropertyID),
			Address:         createPropertyDTO.PropertyDetails.Address,
			ApartmentNumber: createPropertyDTO.PropertyDetails.ApartmentNumber,
			Floors:          createPropertyDTO.PropertyDetails.Floors,
			Rooms:           createPropertyDTO.PropertyDetails.Rooms,
			Area:            createPropertyDTO.PropertyDetails.Area,
		},
	})

	if err != nil {
		return h.errorResponse(c, err)
	}

	return h.successResponse(c, "created")
}

func (h *Handler) updateProperty(c echo.Context) error {
	var updatePropertyDTO dto.UpdatePropertyDTO

	err := c.Bind(&updatePropertyDTO)
	if err != nil {
		return h.errorResponse(c, err)
	}

	_, err = h.propertyService.Update(c.Request().Context(), domain.UpdatePropertyParam{
		Address:         updatePropertyDTO.Address,
		ApartmentNumber: updatePropertyDTO.ApartmentNumber,
		Floors:          updatePropertyDTO.Floors,
		Rooms:           updatePropertyDTO.Rooms,
		Area:            updatePropertyDTO.Area,
	}, domain.ID(updatePropertyDTO.ID))

	if err != nil {
		return h.errorResponse(c, err)
	}

	return h.successResponse(c, "updated")
}
