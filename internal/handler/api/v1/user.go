package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/umed-hotamov/realty-service/internal/domain"
	"github.com/umed-hotamov/realty-service/internal/handler/api/v1/dto"
)

func (h *Handler) InitUserRoutes(api *echo.Group) {
  user := api.Group("/user") 
  {
    userAuth := user.Group("/auth")
    {
      userAuth.POST("/sign-up", h.userSignUp)
      userAuth.POST("/sign-in", h.userSignIn)
      userAuth.POST("/logout",  h.userLogout)
      userAuth.POST("/refresh", h.userRefresh)
    }
  }
}
 
func (h *Handler) userSignUp(c echo.Context) error {
  var signUpDTO dto.SignUpDTO
  
  err := c.Bind(signUpDTO)
  if err != nil {
    return h.errorResponse(c, err)
  }

  err = h.authService.SignUp(c.Request().Context(), domain.SignUpParam{
    Role:     signUpDTO.Role,
    Name:     signUpDTO.Name,    
    Surname:  signUpDTO.Surname,
    Email:    signUpDTO.Email,   
    Password: signUpDTO.Password,
    Phone:    signUpDTO.Phone, 
  })

  if err != nil {
    return h.errorResponse(c, err)
  }

  return h.successResponse(c, "created")
}

func (h *Handler) userSignIn(c echo.Context) error {
	var signInDTO dto.SignInDTO
	err := c.Bind(&signInDTO)
	if err != nil {
		return h.errorResponse(c, err)
	}
	authDetails, err := h.authService.SignIn(c.Request().Context(), domain.SignInParam{
		Email:       signInDTO.Email,
		Password:    signInDTO.Password,
		Fingerprint: signInDTO.Fingerprint,
	})
	if err != nil {
		return h.errorResponse(c, err)
	}

  c.SetCookie(&http.Cookie{
    Name:     "refreshToken",
    Value:    authDetails.AccessToken.String(),
    Path:     "/",
    Domain:   h.config.Server.Host,
    MaxAge:   86400,
    HttpOnly: false,
    Secure:   false,
    SameSite: http.SameSiteLaxMode,
  })

	return h.successResponse(c, authDetails.AccessToken.String())
}

func (h *Handler) userLogout(c echo.Context) error {
	refreshCookie, err := c.Cookie("refreshToken")
	if err != nil {
		return h.errorResponse(c, ErrorUnauthorized)
	}

	err = h.authService.LogOut(c.Request().Context(), domain.Token(refreshCookie.Value))
	if err != nil {
		h.errorResponse(c, err)
	}

  c.SetCookie(&http.Cookie{
    Name:     "refreshToken",
    Value:    "",
    Path:     "/",
    Domain:   h.config.Server.Host,
    MaxAge:   -1,
    HttpOnly: false,
    Secure:   false,
  })

	return h.successResponse(c, "successfully logged out")
}

func (h *Handler) userRefresh(c echo.Context) error {
	var refreshDTO dto.RefreshDTO
	err := c.Bind(&refreshDTO)
	if err != nil {
		return h.errorResponse(c, err)
	}

	refreshCookie, err := c.Cookie("refreshToken")
	if err != nil {
    return h.errorResponse(c, ErrorUnauthorized)
	}

	authDetails, err := h.authService.Refresh(c.Request().Context(), domain.Token(refreshCookie.Value), refreshDTO.Fingerprint)
	if err != nil {
		return h.errorResponse(c, err)
	}

  c.SetCookie(&http.Cookie{
    Name:     "refreshToken",
    Value:    authDetails.AccessToken.String(),
    Path:     "/",
    Domain:   h.config.Server.Host,
    MaxAge:   86400,
    HttpOnly: false,
    Secure:   false,
    SameSite: http.SameSiteLaxMode,
  })

	return h.successResponse(c, authDetails.AccessToken.String())
}
