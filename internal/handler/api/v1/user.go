package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitUserRoutes(api *echo.Group) {
  users := api.Group("/users")
  users.GET("/", h.getAllUsers)
}
 
func (h *Handler) getAllUsers(c echo.Context) error {
  users, err := h.userService.GetAll(c.Request().Context())
  if err != nil {
    return h.errorResponse(c, err)
  }

  usersDTO := make([]dto.UserDTO, len(users))
  for i, user := range users {
    usersDTO[i] = dto.NewUserDTO(user)
  }
  
  return h.successResponse(c, usersDTO)
}
