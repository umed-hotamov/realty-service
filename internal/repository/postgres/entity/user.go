package entity

import (
	"github.com/google/uuid"
	"github.com/umed-hotamov/realty-service/internal/domain"
)

const (
  RoleRealtor       = "realtor"
  RoleAgency        = "agency"
  RolePrivatePerson = "private_person"
)

type PostgresUser struct {
  ID       uuid.UUID `db:"id"`
  Role     string    `db:"role"`
  Name     string    `db:"name"`
  Surname  string    `db:"surname"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
  Phone    string    `db:"phone"`
}

func (u *PostgresUser) ToDomain() domain.User {
  var role domain.UserRole
  
  switch u.Role {
  case RoleRealtor:
    role = domain.Realtor
  case RoleAgency:
    role = domain.Agency
  case RolePrivatePerson:
    role = domain.PrivatePerson
  }

  return domain.User{
    ID:       domain.ID(u.ID.String()),
    Role:     role,
    Name:     u.Name,
    Surname:  u.Surname,
    Email:    u.Email,
    Password: u.Password,
    Phone:    u.Phone,
  }
}

func NewPostgresUser(user domain.User) PostgresUser {
  id, _ := uuid.Parse(user.ID.String())
  
  var role string
  switch user.Role {
  case domain.Realtor:
    role = RoleRealtor
  case domain.Agency:
    role = RoleAgency
  case domain.PrivatePerson:
    role = RolePrivatePerson 
  }

  return PostgresUser{
    ID:       id,
    Role:     role,
    Name:     user.Name,
    Surname:  user.Surname,
    Email:    user.Email,
    Password: user.Password,
    Phone:    user.Phone,
  }
}
