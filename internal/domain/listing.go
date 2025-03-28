package domain

type OfferType int
type ListingStatusType int

const (
	Rent OfferType = iota
	Sale

	Available ListingStatusType = iota
	Sold
	ForRent
	Draft
)

type Listing struct {
	ID          ID
	UserID      ID
	PropertyID  ID
	Title       string
	City        string
	Offer       OfferType
	Price       int
	Description string
	Status      ListingStatusType
}
