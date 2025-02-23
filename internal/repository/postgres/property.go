package repository

// import (
// 	"context"
// 	"database/sql"

// 	"github.com/jmoiron/sqlx"

// 	"github.com/umed-hotamov/realty-service/internal/domain"
// 	"github.com/umed-hotamov/realty-service/internal/repository/postgres/entity"
// )

// type PostgresPropertyRepo struct { 
//   db *sqlx.DB
// }

// func NewPropertyRepo(db *sqlx.DB) *PostgresPropertyRepo {
//   return &PostgresPropertyRepo{
//     db:db,
//   }
// }

// const (
//   PropertySelectAll = "SELECT * FROM public.property"
//   
//   PropertySelectByAddress = "SELECT * FROM public.property p" +
//   "LEFT JOIN public.flat f ON f.property_id = p.id" +  
//   "LEFT JOIN public.private_house h ON h.property_id = p.id" + 
//   "WHERE p.address = 'abc';"
// )

// func (p *PostgresPropertyRepo) GetAll(ctx context.Context) ([]domain.Property, error) {
//   var pgProperties []entity.PostgresProperty
//   err := p.db.SelectContext(ctx, &pgProperties, PropertySelectAll)
//   if err != nil {
//     if err == sql.ErrNoRows {
//       return nil, err
//     }
//     return nil, err
//   }

//   properties := make([]domain.Property, len(pgProperties))
//   for i, p := range pgProperties {
//     properties[i] = p.ToDomain()
//   }

//   return properties, nil
// }

// func (p *PostgresPropertyRepo) GetPropertyByAddress(ctx context.Context, address string) ([]domain.Property, error) {

// }

// func (p *PostgresPropertyRepo) GetApprovedFlatsByHouseID(ctx context.Context, houseID int) ([]domain.Flat, error) {
//   return approvedFlats, nil
// }

// func (p *PostgresPropertyRepo) Create(ctx context.Context, flat domain.Flat) (domain.Flat, error) {
// }

// func (p *PostgresPropertyRepo) Update(ctx context.Context, flatID int) (domain.Flat, error) {
// }
