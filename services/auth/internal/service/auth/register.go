package authsvc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/hashgenerator"
)

func (as *authService) Register(
	ctx context.Context,
	email, login, password string,
) (uuid.UUID, error) {
	const op = "service.auth.Register"

	as.log.Info("attempting register user")

	passwordHash, err := as.hasher.PasswordHash(password)
	if err != nil {
		as.log.Error("failed to generate password hash", slogger.Err(err))

		return uuid.UUID{}, fmt.Errorf("%s: %w", op, err)
	}

	verifyToken, err := hashgenerator.NewHash()
	if err != nil {
		as.log.Error("failed to generate verify token", slogger.Err(err))

		return uuid.UUID{}, fmt.Errorf("%s: %w", op, err)
	}

	user := entity.User{
		Email:        email,
		Login:        login,
		PasswordHash: string(passwordHash),
		UserRole:     entity.SimpleUser,
		VerifyToken:  sql.NullString{String: verifyToken, Valid: true},
	}

	id, err := as.userRepo.Create(ctx, &user)
	if err != nil {
		as.log.Error("failed to create user", slogger.Err(err))

		return uuid.UUID{}, fmt.Errorf("%s: %w", op, AuthErrors(err))
	}

	return id, nil
}
