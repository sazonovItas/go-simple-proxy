package v1

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
	httphandler "github.com/sazonovItas/proxy-manager/services/gateway/internal/handler/http"
)

func (h *handler) Register(c echo.Context) error {
	var u entity.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httphandler.NewError(err))
	}

	if _, err := h.userSvc.Register(c.Request().Context(), u.Email, u.Login, u.Passsword); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httphandler.NewError(err))
	}

	return c.JSON(http.StatusCreated, newRegisterResponse())
}

func (h *handler) Login(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httphandler.NewError(err))
	}

	token, err := h.userSvc.Login(c.Request().Context(), req.User.Login, req.User.Password)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, httphandler.NewError(err))
		case errors.Is(err, adapter.ErrUserInvalidCreditanals):
			return c.JSON(http.StatusForbidden, httphandler.NewError(err))
		}

		return c.JSON(http.StatusInternalServerError, httphandler.NewError(err))
	}

	return c.JSON(http.StatusOK, newLoginResponse(token))
}
