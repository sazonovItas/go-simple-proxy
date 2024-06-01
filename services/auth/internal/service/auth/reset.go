package authsvc

import (
	"context"
	"errors"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/hasher"
)

func (as *authService) GenerateResetToken(ctx context.Context, email string) error {
	const op = "service.auth.GenerateResetToken"

	as.log.Info("attemting create new reset token")

	resetToken, err := hasher.NewRandomHash()
	if err != nil {
		as.log.Error("failed generate reset token", slogger.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	err = as.userRepo.NewResetToken(ctx, email, resetToken)
	if err != nil {
		as.log.Error("failed create new reset token", slogger.Err(err))

		if errors.Is(err, adapter.ErrUserNotFound) {
			return fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (as *authService) ValidateResetToken(ctx context.Context, resetToken string) error {
	const op = "service.auth.ValidateResetToken"

	as.log.Info("attemting validate reset token")

	_, err := as.userRepo.UserByResetToken(ctx, resetToken)
	if err != nil {
		if errors.Is(err, adapter.ErrUserNotFound) {
			return fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (as *authService) ResetPassword(ctx context.Context, resetToken, password string) error {
	const op = "service.auth.ResetPassword"

	as.log.Info("attemting reset password")

	passwordHash, err := as.hasher.PasswordHash(password)
	if err != nil {
		as.log.Error("failed to generate password hash", slogger.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	err = as.userRepo.UpdatePasswordByResetToken(ctx, resetToken, string(passwordHash))
	if err != nil {
		as.log.Error("failed update password by reset token", slogger.Err(err))

		if errors.Is(err, adapter.ErrResetTokenNotFound) {
			return fmt.Errorf("%s: %w", op, ErrResetTokenNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
