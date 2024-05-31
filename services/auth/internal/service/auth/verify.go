package authsvc

import (
	"context"
	"errors"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
)

func (as *authService) VerifyEmail(ctx context.Context, verifyToken string) error {
	const op = "service.auth.VerifyEmail"

	as.log.Info("attemting verify email")

	err := as.userRepo.VerifyEmail(ctx, verifyToken)
	if err != nil {
		as.log.Error("failed to verify email", slogger.Err(err))

		if errors.Is(err, adapter.ErrVerifyTokenNotFound) {
			return fmt.Errorf("%s: %w", op, ErrVerifyTokenNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
