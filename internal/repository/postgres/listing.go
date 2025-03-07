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
  

type PgListingRepo struct {
  db *sqlx.DB
}

func NewPgLisingRepo(db *sqlx.DB) *PgListingRepo {
  return &PgListingRepo{
    db: db,
  }
}

const (
  getAll      = "SELECT * FROM public.listing"
  getByID     = "SELECT * FROM public.listing id = $1" 
  getByUserID = "SELECT * FROM public.listing WHERE user_id = $1"
  deleteByID  = "DELETE FROM public.listing WHERE id = $1"
)


func (l *PgListingRepo) GetAll(ctx context.Context) ([]domain.Listing, error) {
  var pgListings []entity.PgListing

  err := l.db.SelectContext(ctx, &pgListings, getAll)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
  }

  listings := make([]domain.Listing, len(pgListings))
  for i, l := range pgListings {
    listings[i] = l.ToDomain()
  }

  return listings, nil
}

func (l *PgListingRepo) GetUserListings(ctx context.Context, userID int) ([]domain.Listing, error) {
  var pgListings []entity.PgListing

  err := l.db.SelectContext(ctx, &pgListings, getByUserID, userID)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
 }
  listings := make([]domain.Listing, len(pgListings))
  for i, l := range pgListings {
    listings[i] = l.ToDomain()
  }

  return listings, nil
}

func (l *PgListingRepo) GetListingByID(ctx context.Context, id domain.ID) (domain.Listing, error) {
  var pgListing entity.PgListing

  err := l.db.GetContext(ctx, &pgListing, getByUserID, id)
  if err != nil {
    if err == sql.ErrNoRows {
      return domain.Listing{}, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return domain.Listing{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
 }

  return pgListing.ToDomain(), nil
}

func (l *PgListingRepo) Create(ctx context.Context, listing domain.Listing) (domain.Listing, error) {
  pgListing := entity.NewPgListing(listing)
  query := entity.InsertQueryString(pgListing, "listing")

  _, err := l.db.NamedExecContext(ctx, query, pgListing)
  if err != nil {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
      if pgErr.Code == PgUniqueViolationCode {
        return domain.Listing{}, errors.Wrap(errs.ErrDuplicate, err.Error())
      } else {
        return domain.Listing{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
      }
    } else {
      return domain.Listing{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
    }
  }

  return l.GetListingByID(ctx, listing.ID)
}

func (l *PgListingRepo) Update(ctx context.Context, listing domain.Listing) (domain.Listing, error) {
  pgListing := entity.NewPgListing(listing)
  query := entity.UpdateQueryString(pgListing, "listing")

  _, err := l.db.NamedExecContext(ctx, query, pgListing)
  if err != nil {
    return domain.Listing{}, errors.Wrap(errs.ErrNotExist, err.Error())
  }

  return l.GetListingByID(ctx, listing.ID)
}

func (l *PgListingRepo) Delete(ctx context.Context, listingID int) error {
  _, err := l.db.ExecContext(ctx, deleteByID, listingID)
  if err != nil {
    return errors.Wrap(errs.ErrDeleteFailed, err.Error())
  }

  return nil
}
