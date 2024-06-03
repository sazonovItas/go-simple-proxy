package v1

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h *handler) InitRoutes(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initUserRoutes(v1)
		h.initProxyRoutes(v1)
	}
}

func (h *handler) initAuthRoutes(api *echo.Group) {
	auth := api.Group("/users")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}

func (h *handler) initUserRoutes(api *echo.Group) {
	user := api.Group("/user", echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := h.userSvc.ValidateToken(c.Request().Context(), auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	}))
	{
		user.GET("/account", h.CurrentUser)
		user.GET("/request", h.CurrentUserRequests)
	}
}

func (h *handler) initProxyRoutes(api *echo.Group) {
	proxy := api.Group("/proxy", echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: h.parseTokenFunc,
	}))
	{
		proxy.GET("", h.Proxy)
	}
}

func (h *handler) parseTokenFunc(c echo.Context, auth string) (interface{}, error) {
	claims, err := h.userSvc.ValidateToken(c.Request().Context(), auth)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
