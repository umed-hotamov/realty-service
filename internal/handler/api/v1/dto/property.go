package dto

import (
	"github.com/umed-hotamov/realty-service/internal/domain"
)

const (
  Apartment    = "apartment"
  House        = "house"
  
  Cottage      = "cottage"
  CountryHouse = "country_house"
  TownHouse    = "townhome"
  TinyHome     = "tiny_home"
  
  New          = "new"
  Secondary    = "secondary"
)

type PropertyDTO struct {
  ID              string             `json:"id"`
  OwnerID         string             `json:"owner_id"`
  PropertyType    string             `json:"property_type"`
  HouseType       string             `json:"house_type"`
  PropertyDetails PropertyDetailsDTO `json:"property_details"`
  BuildingDetails BuildingDetailsDTO `json:"building_details"`
}

type PropertyDetailsDTO struct {
  PropertyID      string `json:"property_id"`
  Address         string `json:"address"`
  ApartmentNumber int    `json:"apartment_number"`
  Floors          int    `json:"floors"`
  Rooms           int    `json:"rooms"`
  Area            int    `json:"area"`
}

type BuildingDetailsDTO struct {
  BuildingID       string `json:"building_id"`
  Address          string `json:"address"`
  Developer        string `json:"developer"`
  Floors           int    `json:"floors"`
  ConstructionYear int    `json:"construction_year"`
  ParkingPlace     bool   `json:"parking_place"`
}

type CreatePropertyDTO struct {
  OwnerID         string              `json:"owner_id"`
  BuildingID      *string             `json:"building_id"`
  PropertyType    domain.PropertyType `json:"property_type"`
  HouseType       *domain.HouseType   `json:"house_type,omitempty"`
  PropertyDetails PropertyDetailsDTO  `json:"property_details"`
}

type UpdatePropertyDTO struct {
  ID              string `json:"id"`
  Address         string `json:"address"`
  ApartmentNumber int    `json:"apartment_number"`
  Floors          int    `json:"floors"`
  Rooms           int    `json:"rooms"`
  Area            int    `json:"area"`
}

func NewPropertyDTO(property domain.Property) PropertyDTO {
  var (
    propertyType string
    houseType    string
  )

  switch property.PropertyType {
  case domain.Apartment:
    propertyType = Apartment
  case domain.House:
    propertyType = House
  }
  
  switch property.HouseType {  
  case domain.Cottage: 
    houseType = Cottage     
  case domain.CountryHouse:
    houseType = CountryHouse
  case domain.TownHouse:
    houseType = TownHouse   
  case domain.TinyHome: 
    houseType = TinyHome
  }
  
  return PropertyDTO{
    ID:              property.ID.String(), 
    OwnerID:         property.OwnerID.String(),
    PropertyType:    propertyType,
    HouseType:       houseType,
    PropertyDetails: PropertyDetailsDTO{
      PropertyID:      property.ID.String(),
      Address:         property.PropertyDetails.Address,         
      ApartmentNumber: property.PropertyDetails.ApartmentNumber,
      Floors:          property.PropertyDetails.Floors,
      Rooms:           property.PropertyDetails.Rooms,
      Area:            property.PropertyDetails.Area,
    },
    BuildingDetails: BuildingDetailsDTO{
      BuildingID:       property.BuildingID.String(),      
      Address:          property.BuildingDetails.Address,            
      Developer:        property.BuildingDetails.Developer,
      Floors:           property.BuildingDetails.Floors,
      ConstructionYear: property.BuildingDetails.ConstructionYear,
      ParkingPlace:     property.BuildingDetails.ParkingPlace, 
    },
  }
}
