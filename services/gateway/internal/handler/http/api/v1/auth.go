package v1

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Register(c echo.Context) error {
	return nil
}

func (h *handler) Login(c echo.Context) error {
	return nil
}

func (h *handler) initAuthRoutes(api *echo.Group) {
	auth := api.Group("")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", nil)
	}
}
