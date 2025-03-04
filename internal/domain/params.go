package domain

type CreateUserParam struct {
  Role     UserRole
  Name     string
  Surname  string
  Email    string
  Password string
  Phone    string
}

type UpdateUserParam struct {
  Name    string
  Surname string
  Phone   string
}
