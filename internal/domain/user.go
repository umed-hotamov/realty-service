package domain

type UserRole int

const (
  Seller UserRole = iota
  Moderator
)

type User struct {
  ID       ID
  Email    string
  Password string
  Phone    string
  Role     UserRole
}
