package domain

type PropertyUnit interface{}

type Property struct {
  ID        int
  UserID    int
  Type      bool
  OfferType bool
  Price     int
  Status    int
  Unit      PropertyUnit
}

type PropertyFilter struct {
  Type     bool
  MinPrice int
  MaxPrice int
  Address  string
} 

type Flat struct {
  ID         int
  HouseID    int
  PropertyID int
  Number     int
  Price      int
  Rooms      int
  Square     int
}

type PrivateHouse struct {
  ID         int
  PropertyID int
  Address    string
  Rooms      int
  Square     int
}
