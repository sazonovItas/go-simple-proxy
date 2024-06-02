package v1

import "github.com/labstack/echo/v4"

func (h *handler) Proxy(c echo.Context) error {
	return nil
}

func (h *handler) initProxyRoutes(api *echo.Group) {
	proxy := api.Group("/proxy")
	{
		proxy.GET("", h.Proxy)
	}
}
