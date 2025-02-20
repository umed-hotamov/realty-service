package entity

// import (
// 	"github.com/google/uuid"
// 	"github.com/umed-hotamov/realty-service/internal/domain"
// )

// const (
//   RentType         = "rent"
//   SaleType         = "sale"
//   
//   FlatType         = "flat"
//   PrivateHouseType = "private_house"
//   
//   SoldType         = "sold"
//   ForRentType      = "for_rent"
//   AvailableType        = "alive"
//   
//   CreatedType      = "created"
//   DeclinedType      = "declined"
//   ApprovedType     = "approved"
//   OnModerationType = "on_moderation"
// )

// type PostgresProperty struct {
//   ID      uuid.UUID `db:"id"`
//   OwnerID uuid.UUID `db:"owner_id"`
//   Type    string    `db:"type"`
//   Offer   string    `db:"offer"`
//   PStatus string    `db:"p_status"`
//   MStatus string    `db:"m_status"`
//   Price   int64     `db:"price"`
//   Rooms   int       `db:"rooms"`
//   Square  int       `sb:"square"`
// }

// func (p *PostgresProperty) ToDomain() domain.Property {
//   var (
//     propertyType     domain.PropertyType
//     offerType        domain.OfferType
//     propertyStatus   domain.PropertyStatus
//     moderationStatus domain.ModerationStatus
//   )
//   
//   switch p.Type {
//   case FlatType:
//     propertyType = domain.FlatType
//   case PrivateHouseType:
//     propertyType = domain.PrivateHouseType
//   }

//   switch p.Offer {
//   case RentType:
//     offerType = domain.Rent
//   case SaleType:
//     offerType = domain.Sale
//   }

//   switch p.PStatus {
//   case SoldType:
//     propertyStatus = domain.Sold
//   case ForRentType:
//     propertyStatus = domain.ForRent
//   case AvailableType:
//     propertyStatus = domain.Available
//   }
//   
//   switch p.MStatus{
//   case CreatedType:
//     moderationStatus = domain.Created
//   case DeclinedType:
//     moderationStatus = domain.Declined
//   case ApprovedType:
//     moderationStatus = domain.Approved
//   case OnModerationType:
//     moderationStatus = domain.OnModeration
//   }

//   return domain.Property{
//     ID:      domain.ID(p.ID.String()),
//     OwnerID: domain.ID(p.OwnerID.String()),
//     Type:    propertyType,
//     Offer:   offerType,
//     PStatus: propertyStatus,
//     MStatus: moderationStatus,
//     Price:   p.Price,
//     Rooms:   p.Rooms,
//     Square:  p.Square,
//   }
// }
