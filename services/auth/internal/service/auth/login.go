package authsvc

import (
	"context"
	"errors"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/jwt"
)

// TODO: check user email verification
func (as *authService) Login(
	ctx context.Context,
	login, password string,
) (string, error) {
	const op = "service.auth.Login"

	as.log.Info("attemting to login user")

	user, err := as.userRepo.UserByLogin(ctx, login)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrUserNotFound):
			as.log.Warn("user not found", slogger.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrUserNotFound)
		default:
			as.log.Error("failed to get user", slogger.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
		}
	}

	if err = as.hasher.Compare(user.PasswordHash, password); err != nil {
		as.log.Info("invalid credentials")

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	userInfo := entity.UserInfo{
		ID:    user.ID.String(),
		Email: user.Email,
		Login: user.Login,
		Role:  user.UserRole,
	}

	token, err := jwt.NewToken(userInfo, as.authSecret, as.tokenTTL)
	if err != nil {
		as.log.Error("failed to generate jwt token", err)

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}
