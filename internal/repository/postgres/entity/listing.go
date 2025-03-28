package entity

import (
	"github.com/google/uuid"
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

type PgListing struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	PropertyID  uuid.UUID `db:"property_id"`
	Title       string    `db:"title"`
	City        string    `db:"city"`
	Offer       string    `db:"offer"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
}

func (l *PgListing) ToDomain() domain.Listing {
	var (
		offerType  domain.OfferType
		statusType domain.ListingStatusType
	)

	switch l.Offer {
	case Rent:
		offerType = domain.Rent
	case Sale:
		offerType = domain.Sale
	}

	switch l.Status {
	case Available:
		statusType = domain.Available
	case Sold:
		statusType = domain.Sold
	case ForRent:
		statusType = domain.ForRent
	case Draft:
		statusType = domain.Draft
	}

	return domain.Listing{
		ID:          domain.ID(l.ID.String()),
		UserID:      domain.ID(l.UserID.String()),
		PropertyID:  domain.ID(l.PropertyID.String()),
		Title:       l.Title,
		City:        l.City,
		Offer:       offerType,
		Price:       l.Price,
		Description: l.Description,
		Status:      statusType,
	}
}

func NewPgListing(listing domain.Listing) PgListing {
	listingID, _ := uuid.Parse(listing.ID.String())
	userID, _ := uuid.Parse(listing.UserID.String())
	propertyID, _ := uuid.Parse(listing.PropertyID.String())

	var (
		offer  string
		status string
	)

	switch listing.Offer {
	case domain.Rent:
		offer = Rent
	case domain.Sale:
		offer = Sale
	}

	switch listing.Status {
	case domain.Available:
		status = Available
	case domain.Sold:
		status = Sold
	case domain.ForRent:
		status = ForRent
	case domain.Draft:
		status = Draft
	}

	return PgListing{
		ID:          listingID,
		UserID:      userID,
		PropertyID:  propertyID,
		Title:       listing.Title,
		City:        listing.City,
		Offer:       offer,
		Price:       listing.Price,
		Description: listing.Description,
		Status:      status,
	}
}
