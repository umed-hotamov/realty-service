package domain

type UserRole int

const (
	Realtor UserRole = iota
	Agency
	PrivatePerson
)

type User struct {
	ID       ID
	Role     UserRole
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    string
}
