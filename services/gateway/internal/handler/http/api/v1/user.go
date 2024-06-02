package v1

import "github.com/labstack/echo/v4"

func (h *handler) Account(c echo.Context) error {
	return nil
}

func (h *handler) Request(c echo.Context) error {
	return nil
}

func (h *handler) initUserRoutes(api *echo.Group) {
	user := api.Group("/user")
	{
		user.GET("/account", nil)
		user.GET("/request", nil)
	}
}
