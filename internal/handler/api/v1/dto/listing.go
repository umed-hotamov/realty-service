package dto

import (
	"github.com/umed-hotamov/realty-service/internal/domain"
)

const (
	Rent      = "rent"
	Sale      = "sale"
	Available = "available"
	Sold      = "sold"
	ForRent   = "for_rent"
	Draft     = "draft"
)

type ListingDTO struct {
	UserID      string `json:"user_id"`
	PropertyID  string `json:"property_id"`
	Title       string `json:"title"`
	City        string `json:"city"`
	Offer       string `json:"offer"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateListingDTO struct {
	PropertyID  string                   `json:"property_id"`
	Title       string                   `json:"title"`
	City        string                   `json:"city"`
	Offer       domain.OfferType         `json:"offer"`
	Price       int                      `json:"price"`
	Description string                   `json:"description"`
	Status      domain.ListingStatusType `json:"status"`
}

type UpdateListingDTO struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	City        string           `json:"city"`
	Offer       domain.OfferType `json:"offer"`
	Price       int              `json:"price"`
	Description string           `json:"description"`
}

func NewListingDTO(listing domain.Listing) ListingDTO {
	var (
		offerType  string
		statusType string
	)

	switch listing.Offer {
	case domain.Rent:
		offerType = Rent
	case domain.Sale:
		offerType = Sale
	}

	switch listing.Status {
	case domain.Available:
		statusType = Available
	case domain.Sold:
		statusType = Sold
	case domain.ForRent:
		statusType = ForRent
	case domain.Draft:
		statusType = Draft
	}

	return ListingDTO{
		UserID:      listing.UserID.String(),
		PropertyID:  listing.PropertyID.String(),
		Title:       listing.Title,
		City:        listing.City,
		Offer:       offerType,
		Price:       listing.Price,
		Description: listing.Description,
		Status:      statusType,
	}
}
