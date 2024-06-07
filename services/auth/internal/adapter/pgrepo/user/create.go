package pguser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	authsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/auth"
)

func (us *userRepository) NewUser(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	const op = "adapter.pgrepo.user.Create"

	const query = `INSERT INTO %s (id, email, login, password_hash, user_role)
	VALUES (:id, :email, :login, :password_hash, :user_role)`

	const checkUserQuery = "SELECT * FROM %s WHERE email=$1 OR login=$2"

	tx, err := us.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return uuid.UUID{}, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	checkUserStmt, err := tx.PreparexContext(ctx, us.table(checkUserQuery))
	if err != nil {
		return uuid.UUID{}, err
	}
	defer checkUserStmt.Close()

	var checkUser entity.User
	err = checkUserStmt.GetContext(ctx, &checkUser, user.Email, user.Login)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		if err == nil {
			if user.Email == checkUser.Email {
				return uuid.UUID{}, fmt.Errorf("%s: %w", op, authsvc.ErrUserWithEmailAlreadyExists)
			}

			if user.Login == checkUser.Login {
				return uuid.UUID{}, fmt.Errorf("%s: %w", op, authsvc.ErrUserWithLoginAlreadyExists)
			}
		}

		return uuid.UUID{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := tx.PrepareNamedContext(ctx, us.table(query))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: failed to prepare statement %w", op, err)
	}
	defer stmt.Close()

	user.ID = uuid.New()
	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: failed create user %w", op, err)
	}

	if err = tx.Commit(); err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: failed to commit transaction: %w", op, err)
	}

	return user.ID, nil
}
