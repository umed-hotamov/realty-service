package dto

import "github.com/umed-hotamov/realty-service/internal/domain"

type SignUpDTO struct {
  ID       string           `json:"id"`
  Role     domain.UserRole  `json:"role"`
  Name     string           `json:"name"`
  Surname  string           `json:"surname"`
  Email    string           `json:"email"`
  Password string           `json:"password"`
  Phone    string           `json:"phone"`
}

type SignInDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fingerprint string `json:"fingerprint"`
}

type RefreshDTO struct {
	Fingerprint string `json:"fingerprint"`
}
