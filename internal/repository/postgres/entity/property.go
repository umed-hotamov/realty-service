package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/umed-hotamov/realty-service/internal/domain"
)

const (
  TypeApartment    = "apartment"
  TypeHouse        = "house"
  
  TypeCottage      = "cottage"
  TypeCountryHouse = "country_house"
  TypeTownHouse    = "townhome"
  TypeTinyHome     = "tiny_home"
  
  TypeNew          = "new"
  TypeSecondary    = "secondary"
)

type PgProperty struct {
  ID              uuid.UUID     `db:"id"`
  OwnerID         uuid.UUID     `db:"owner_id"`
  BuildingID      uuid.NullUUID `db:"building_id"`
  PropertyType    string        `db:"property_type"`
  HouseType       null.String   `db:"house_type"`
}

type PgPropertyDetails struct {
  PropertyID      uuid.UUID `db:"property_id"`
  Address         string    `db:"address"`
  ApartmentNumber int       `db:"apartment_number"`
  Floors          int       `db:"floors"`
  Rooms           int       `db:"rooms"`
  Area            int       `db:"area"`
}

type PgApartmentBuilding struct {
  ID   uuid.UUID `db:"id"`
  Type string    `db:"type"` 
}

type PgBuildingDetails struct {
  BuildingID       uuid.UUID `db:"building_id"`
  Address          string    `db:"address"`
  Developer        string    `db:"developer"`
  Floors           int       `db:"floors"`
  ConstructionYear int       `db:"construction_year"`
  ParkingPlace     bool      `db:"parking_place"`
}

func (p *PgProperty) ToDomain(pDetails PgPropertyDetails, bDetails PgBuildingDetails, apartmentBuilding PgApartmentBuilding) domain.Property {
  var (
    propertyType domain.PropertyType
    houseType    domain.HouseType
    buildingType domain.BuildingType
  )

  switch p.PropertyType {
  case TypeApartment:
    propertyType = domain.Apartment
  case TypeHouse:
    propertyType = domain.House
  }
  
  switch p.HouseType.String {
  case TypeCottage:
    houseType = domain.Cottage
  case TypeCountryHouse:
    houseType = domain.CountryHouse
  case TypeTownHouse:
    houseType = domain.TownHouse
  case TypeTinyHome:
    houseType = domain.TinyHome
  }

  switch apartmentBuilding.Type {
  case TypeNew:
    buildingType = domain.New
  case TypeSecondary:
    buildingType = domain.Secondary
  }
  
  propertyDetails := domain.PropertyDetails{
    PropertyID:      domain.ID(pDetails.PropertyID.String()),      
    Address:         pDetails.Address,
    ApartmentNumber: pDetails.ApartmentNumber,
    Floors:          pDetails.Floors,          
    Rooms:           pDetails.Rooms,           
    Area:            pDetails.Area,
  }
  
  var buildingDetails domain.BuildingDetails
  if bDetails != (PgBuildingDetails{}) {
    buildingDetails = domain.BuildingDetails{
      BuildingType:     buildingType,
      Address:          bDetails.Address,
      Developer:        bDetails.Developer,
      Floors:           bDetails.Floors,
      ConstructionYear: bDetails.ConstructionYear,
      ParkingPlace:     bDetails.ParkingPlace,
    }
  }

  return domain.Property{
    ID:              domain.ID(p.ID.String()),            
    OwnerID:         domain.ID(p.OwnerID.String()),
    BuildingID:       domain.ID(bDetails.BuildingID.String()),   
    PropertyType:    propertyType,
    HouseType:       houseType,     
    PropertyDetails: propertyDetails,
    BuildingDetails: buildingDetails,
  }
}

func NewPgProperty(property domain.Property) PgProperty {
  propertyID, _ := uuid.Parse(property.ID.String())
  ownerID, _ := uuid.Parse(property.OwnerID.String())
  buildingID, _ := uuid.Parse(property.BuildingID.String())

  var (
    propertyType string
    houseType    string
  )

  switch property.PropertyType {
  case domain.Apartment:
    propertyType = TypeApartment
  case domain.House:
    propertyType = TypeHouse
  }
  
  switch property.HouseType {  
  case domain.Cottage: 
    houseType = TypeCottage     
  case domain.CountryHouse:
    houseType = TypeCountryHouse
  case domain.TownHouse:
    houseType = TypeTownHouse   
  case domain.TinyHome: 
    houseType = TypeTinyHome
  }
  
  return PgProperty{
    ID:           propertyID, 
    OwnerID:      ownerID,
    BuildingID:   uuid.NullUUID{
      UUID:  buildingID,
      Valid: true,
    },
    PropertyType: propertyType,
    HouseType:    null.StringFrom(houseType),
  }
}

func NewPgPropertyDetails(details domain.PropertyDetails) PgPropertyDetails {
  propertyID, _ := uuid.Parse(details.PropertyID.String())
  
  return PgPropertyDetails{
    PropertyID:      propertyID,     
    Address:         details.Address,
    ApartmentNumber: details.ApartmentNumber,
    Floors:          details.Floors,
    Rooms:           details.Rooms,
    Area:            details.Area,
  }
}
