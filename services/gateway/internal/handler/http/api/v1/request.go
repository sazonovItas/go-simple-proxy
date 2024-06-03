package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

type userRegisterRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *entity.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	u.Email = r.User.Email
	u.Login = r.User.Login
	u.Passsword = r.User.Password
	return nil
}

type userLoginRequest struct {
	User struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}
