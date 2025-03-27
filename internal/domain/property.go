package domain

type PropertyType     int
type HouseType        int
type BuildingType     int 

const (
  Apartment    PropertyType = iota
  House
  
  Cottage      HouseType = iota
  CountryHouse
  TownHouse
  TinyHome

  New          BuildingType = iota
  Secondary 
)

type Property struct {
  ID              ID
  OwnerID         ID
  BuildingID      ID
  PropertyType    PropertyType
  HouseType       HouseType
  PropertyDetails PropertyDetails
  BuildingDetails BuildingDetails
}

type PropertyDetails struct {
  PropertyID      ID
  Address         string
  ApartmentNumber int
  Floors          int
  Rooms           int
  Area            int
}

type BuildingDetails struct {
  BuildingType     BuildingType
  Address          string
  Developer        string
  Floors           int
  ConstructionYear int
  ParkingPlace     bool
}
