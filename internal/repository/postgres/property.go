package repository

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/errs"
	"github.com/umed-hotamov/realty-service/internal/repository/postgres/entity"
)

type PgPropertyRepo struct {
	db *sqlx.DB
}

func NewPropertyRepo(db *sqlx.DB) *PgPropertyRepo {
	return &PgPropertyRepo{
		db: db,
	}
}

const (
	PropertySelectAll     = "SELECT * FROM public.property"
	SelectPropertyDetails = "SELECT * FROM public.property_details WHERE property_id = $1"
	SelectApartment       = "SELECT * FROM public.apartment_building WHERE id = $1"
	SelectUserProperties  = "SELECT * FROM public.property WHERE owner_id = $1"
	SelectPropertyByID    = "SELECT * FROM public.property WHERE id = $1"
	PropertyDelete        = "DELETE FROM public.property WHERE id = $1"
	SelectBuildingDetails = "SELECT * FROM public.building_details WHERE building_id = $1"
)

func (p *PgPropertyRepo) GetAll(ctx context.Context) ([]domain.Property, error) {
	var (
		pgProperties        []entity.PgProperty
		pgPropertyDetails   entity.PgPropertyDetails
		pgApartmentBuilding entity.PgApartmentBuilding
		pgBuildingDetails   entity.PgBuildingDetails
	)

	err := p.db.SelectContext(ctx, &pgProperties, PropertySelectAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		}
		return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
	}

	properties := make([]domain.Property, len(pgProperties))
	for i, property := range pgProperties {
		err = p.db.GetContext(ctx, &pgPropertyDetails, SelectPropertyDetails, property.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.Wrap(errs.ErrNotExist, err.Error())
			}
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}

		if property.BuildingID.Valid {
			err = p.db.GetContext(ctx, &pgApartmentBuilding, SelectApartment, property.BuildingID)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.Wrap(errs.ErrNotExist, err.Error())
				}
				return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}

			err = p.db.GetContext(ctx, &pgBuildingDetails, SelectBuildingDetails, property.BuildingID)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.Wrap(errs.ErrNotExist, err.Error())
				}
				return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		}

		properties[i] = property.ToDomain(pgPropertyDetails, pgBuildingDetails, pgApartmentBuilding)
	}

	return properties, nil
}

func (p *PgPropertyRepo) GetPropertyByID(ctx context.Context, propertyID domain.ID) (domain.Property, error) {
	var (
		pgProperty          entity.PgProperty
		pgPropertyDetails   entity.PgPropertyDetails
		pgApartmentBuilding entity.PgApartmentBuilding
		pgBuildingDetails   entity.PgBuildingDetails
	)

	err := p.db.GetContext(ctx, &pgProperty, SelectPropertyByID, propertyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Property{}, errors.Wrap(errs.ErrNotExist, err.Error())
		}
		return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
	}

	err = p.db.GetContext(ctx, &pgPropertyDetails, SelectPropertyDetails, pgProperty.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Property{}, errors.Wrap(errs.ErrNotExist, err.Error())
		}
		return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
	}

	if pgProperty.BuildingID.Valid {
		err = p.db.GetContext(ctx, &pgApartmentBuilding, SelectApartment, pgProperty.BuildingID)
		if err != nil {
			if err == sql.ErrNoRows {
				return domain.Property{}, errors.Wrap(errs.ErrNotExist, err.Error())
			}
			return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}

		err = p.db.GetContext(ctx, &pgBuildingDetails, SelectBuildingDetails, pgProperty.BuildingID)
		if err != nil {
			if err == sql.ErrNoRows {
				return domain.Property{}, errors.Wrap(errs.ErrNotExist, err.Error())
			}
			return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	return pgProperty.ToDomain(pgPropertyDetails, pgBuildingDetails, pgApartmentBuilding), nil
}

func (p *PgPropertyRepo) GetUserProperties(ctx context.Context, userID domain.ID) ([]domain.Property, error) {
	var (
		pgProperties        []entity.PgProperty
		pgPropertyDetails   entity.PgPropertyDetails
		pgApartmentBuilding entity.PgApartmentBuilding
		pgBuildingDetails   entity.PgBuildingDetails
	)

	err := p.db.SelectContext(ctx, &pgProperties, SelectUserProperties, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		}
		return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
	}

	properties := make([]domain.Property, len(pgProperties))
	for i, property := range pgProperties {
		err = p.db.GetContext(ctx, &pgPropertyDetails, SelectPropertyDetails, property.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.Wrap(errs.ErrNotExist, err.Error())
			}
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}

		if property.BuildingID.Valid {
			err = p.db.GetContext(ctx, &pgApartmentBuilding, SelectApartment, property.BuildingID)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.Wrap(errs.ErrNotExist, err.Error())
				}
				return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}

			err = p.db.GetContext(ctx, &pgBuildingDetails, SelectBuildingDetails, property.BuildingID)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, errors.Wrap(errs.ErrNotExist, err.Error())
				}
				return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		}

		properties[i] = property.ToDomain(pgPropertyDetails, pgBuildingDetails, pgApartmentBuilding)
	}

	return properties, nil
}

func (p *PgPropertyRepo) Create(ctx context.Context, property domain.Property) (domain.Property, error) {
	pgProperty := entity.NewPgProperty(property)
	query := entity.InsertQueryString(pgProperty, "property")

	_, err := p.db.NamedExecContext(ctx, query, pgProperty)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Property{}, errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	pgPropertyDetails := entity.NewPgPropertyDetails(property.PropertyDetails)
	query = entity.InsertQueryString(pgPropertyDetails, "property")

	_, err = p.db.NamedExecContext(ctx, query, pgPropertyDetails)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Property{}, errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Property{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	return p.GetPropertyByID(ctx, property.ID)
}

func (p *PgPropertyRepo) Update(ctx context.Context, property domain.Property) (domain.Property, error) {
	pgProperty := entity.NewPgProperty(property)
	query := entity.UpdateQueryString(pgProperty, "user")

	_, err := p.db.NamedExecContext(ctx, query, pgProperty)
	if err != nil {
		return domain.Property{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
	}

	return p.GetPropertyByID(ctx, property.ID)
}

func (p *PgPropertyRepo) Delete(ctx context.Context, propertyID domain.ID) error {
	_, err := p.db.ExecContext(ctx, PropertyDelete, propertyID)
	if err != nil {
		return errors.Wrap(errs.ErrDeleteFailed, err.Error())
	}

	return nil
}
