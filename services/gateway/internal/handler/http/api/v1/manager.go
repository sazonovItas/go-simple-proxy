package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	httphandler "github.com/sazonovItas/proxy-manager/services/gateway/internal/handler/http"
)

func (h *handler) Proxy(c echo.Context) error {
	proxies, err := h.proxySvc.ProxyInfo(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httphandler.NewError(err))
	}

	return c.JSON(http.StatusOK, newProxyInfoReqsponse(proxies))
}
