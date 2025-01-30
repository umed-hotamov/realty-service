package dto

import "github.com/umed-hotamov/realty-service/internal/domain"

const (
  RoleSeller    = "Seller"
  RoleModerator = "Moderator"
)

type UserDTO struct {
  ID       string `json:"id"`
  Email    string `json:"email"`
  Password string `json:"password"`
  Phone    string `json:"phone"`
  Role     string `json:"role"`
}

func NewUserDTO(user domain.User) UserDTO {
  var role string
  switch user.Role {
  case domain.Seller:
    role = RoleSeller
  case domain.Moderator:
    role = RoleModerator
  }

  return UserDTO{
    ID: user.ID.String(),
    Email: user.Email,
    Password: user.Password,
    Phone: user.Phone,
    Role: role,
  }
}
