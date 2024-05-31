package authsvc

import (
	"context"
	"errors"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/hashgenerator"
)

func (as *authService) ResetToken(ctx context.Context, email string) error {
	const op = "service.auth.ResetToken"

	as.log.Info("attemting create new reset token")

	resetToken, err := hashgenerator.NewHash()
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

func (as *authService) UpdatePassword(ctx context.Context, password, resetToken string) error {
	const op = "service.auth.UpdatePassword"

	as.log.Info("attemting update password")

	passwordHash, err := as.hasher.PasswordHash(password)
	if err != nil {
		as.log.Error("failed to generate password hash", slogger.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	err = as.userRepo.UpdatePasswordByResetToken(ctx, string(passwordHash), resetToken)
	if err != nil {
		as.log.Error("failed update password by reset token", slogger.Err(err))

		if errors.Is(err, adapter.ErrResetTokenNotFound) {
			return fmt.Errorf("%s: %w", op, ErrResetTokenNotFound)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
