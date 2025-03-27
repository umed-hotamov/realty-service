package domain

type CreateUserParam struct {
  Role     UserRole
  Name     string
  Surname  string
  Email    string
  Password string
  Phone    string
}

type UpdateUserParam struct {
  Name    string
  Surname string
  Phone   string
}

type CreatePropertyParam struct {
  OwnerID          ID   
  BuildingID       ID
  PropertyType     PropertyType
  HouseType        HouseType
  PropertyDetails  PropertyDetails
}

type UpdatePropertyParam struct {
  Address         string
  ApartmentNumber int
  Floors          int
  Rooms           int
  Area            int
}

type CreateListingParam struct {
  UserID      ID
  PropertyID  ID
  Title       string
  City        string
  Offer       OfferType
  Price       int
  Description string
  Status      ListingStatusType
}

type UpdateListingParam struct {
  Title       string
  City        string
  Offer       OfferType
  Price       int
  Description string
}

type SignInParam struct {
	Email       string
	Password    string
	Fingerprint string
}

type SignUpParam struct {
	Role     UserRole
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    string
}
