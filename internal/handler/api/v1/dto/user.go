package dto

import "github.com/umed-hotamov/realty-service/internal/domain"

const (
	RoleRealtor       = "realtor"
	RoleAgency        = "agency"
	RolePrivatePerson = "private_person"
)

type UserDTO struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func NewUserDTO(user domain.User) UserDTO {
	var role string

	switch user.Role {
	case domain.Realtor:
		role = RoleRealtor
	case domain.Agency:
		role = RoleAgency
	case domain.PrivatePerson:
		role = RolePrivatePerson
	}

	return UserDTO{
		ID:       user.ID.String(),
		Role:     role,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
}
