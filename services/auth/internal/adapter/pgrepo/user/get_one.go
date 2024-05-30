package pguser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

func (us *userRepository) UserById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	const op = "adapter.pgrepo.user.UserById"

	const query = "SELECT * FROM %s WHERE id=$1"

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement %w", op, err)
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.GetContext(ctx, &user, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, adapter.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: failed get user %w", op, err)
	}

	return &user, nil
}

func (us *userRepository) UserByEmail(ctx context.Context, email string) (*entity.User, error) {
	const op = "adapter.pgrepo.user.UserByEmail"

	const query = "SELECT * FROM %s WHERE email=$1"

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement %w", op, err)
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.GetContext(ctx, &user, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, adapter.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: failed get user %w", op, err)
	}

	return &user, nil
}

func (us *userRepository) UserByLogin(
	ctx context.Context,
	login string,
) (*entity.User, error) {
	const op = "adapter.pgrepo.user.UserByLogin"

	const query = "SELECT * FROM %s WHERE login=$1"

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement %w", op, err)
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.GetContext(ctx, &user, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, adapter.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: failed get user %w", op, err)
	}

	return &user, nil
}
