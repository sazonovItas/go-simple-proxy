package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
	httphandler "github.com/sazonovItas/proxy-manager/services/gateway/internal/handler/http"
)

func (h *handler) CurrentUser(c echo.Context) error {
	id := userIdFromToken(c)

	u, err := h.userSvc.User(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, httphandler.NewError(err))
		}

		return c.JSON(http.StatusInternalServerError, httphandler.NewError(err))
	}

	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *handler) CurrentUserRequests(c echo.Context) error {
	id := userIdFromToken(c)

	requests, err := h.reqSvc.TimestampAndUserId(
		c.Request().Context(),
		time.Now().AddDate(0, 0, -3),
		time.Now(),
		id,
	)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return c.JSON(http.StatusNotFound, httphandler.NewError(err))
		}

		return c.JSON(http.StatusInternalServerError, httphandler.NewError(err))
	}

	return c.JSON(http.StatusOK, newUserProxyRequestsReqsponse(requests))
}

func userIdFromToken(c echo.Context) string {
	token, ok := c.Get("user").(*entity.Token)
	if !ok {
		return ""
	}

	return token.ID
}
