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

type PostgresUserRepo struct {
  db *sqlx.DB
} 

func NewUserRepo(db *sqlx.DB) *PostgresUserRepo {
  return &PostgresUserRepo{
    db:db,
  }
}

const (
  UserSelectAll     = "SELECT * FROM public.user"
  UserSelectByID    = "SELECT * FROM public.user WHERE id = $1"
  UserSelectByRole  = "SELECT * FROM public.user WHERE role = $1"
  UserSelectByEmail = "SELECT * FROM public.user WHERE email = $1"
  UserDelete        = "DELETE FROM public.user WHERE id = $1"
)

func (u *PostgresUserRepo) GetAll(ctx context.Context) ([]domain.User, error) {
  var pgUsers []entity.PostgresUser
  
  err := u.db.SelectContext(ctx, &pgUsers, UserSelectAll)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
  }

  users := make([]domain.User, len(pgUsers))
  for i, user := range pgUsers {
    users[i] = user.ToDomain() 
  }

  return users, nil
}

func (u *PostgresUserRepo) GetByID(ctx context.Context, id domain.ID) (domain.User, error) {
  var pgUser entity.PostgresUser

  err := u.db.GetContext(ctx, &pgUser, UserSelectByID, id)
  if err != nil {
    if err == sql.ErrNoRows {
      return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
  }

  return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) GetByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error) {
  var pgUsers []entity.PostgresUser

  err := u.db.SelectContext(ctx, &pgUsers, UserSelectByRole)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
  }

  users := make([]domain.User, len(pgUsers))
  for i, user := range pgUsers {
    users[i] = user.ToDomain()
  }

  return users, nil 
}

func (u *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
  var pgUser entity.PostgresUser

  err := u.db.GetContext(ctx, &pgUser, UserSelectByEmail, email)
  if err != nil {
    if err == sql.ErrNoRows {
      return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
    }
    return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
  }

  return pgUser.ToDomain(), nil 
}

func (u *PostgresUserRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
  pgUser := entity.NewPostgresUser(user)
  query := entity.InsertQueryString(pgUser, "user")

  _, err := u.db.NamedExecContext(ctx, query, pgUser)
  if err != nil {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
      if pgErr.Code == PgUniqueViolationCode {
        return domain.User{}, errors.Wrap(errs.ErrDuplicate, err.Error())
      } else {
        return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
      }
    } else {
      return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
    }
  }

  return u.GetByID(ctx, user.ID)
}

func (u *PostgresUserRepo) Update(ctx context.Context, user domain.User) (domain.User, error) {
  pgUser := entity.NewPostgresUser(user)
  query := entity.UpdateQueryString(pgUser, "user")

  _, err := u.db.NamedExecContext(ctx, query, pgUser)
  if err != nil {
    return domain.User{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
  }

  return u.GetByID(ctx, user.ID)
}
  
func (u *PostgresUserRepo) Delete(ctx context.Context, userID domain.ID) error {
  _, err := u.db.ExecContext(ctx, UserDelete, userID)
  if err != nil {
    return errors.Wrap(errs.ErrDeleteFailed, err.Error())
  }

  return nil 
}
