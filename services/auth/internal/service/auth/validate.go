package authsvc

import (
	"context"
	"errors"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/jwt"
)

func (as *authService) Validate(ctx context.Context, token string) (entity.UserInfo, error) {
	const op = "service.auth.Validate"

	as.log.Info("attemting validate token")

	claims, err := jwt.ValidateToken(token, as.authSecret)
	if err != nil {
		as.log.Warn("token validation error", slogger.Err(err))

		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return entity.UserInfo{}, fmt.Errorf("%s: %w", op, ErrTokenExpired)
		case errors.Is(err, jwt.ErrTokenMalformed):
			return entity.UserInfo{}, fmt.Errorf("%s, %w", op, ErrTokenMalformed)
		}

		as.log.Error("failed validate token", slogger.Err(err))

		return entity.UserInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	return claims, nil
}
