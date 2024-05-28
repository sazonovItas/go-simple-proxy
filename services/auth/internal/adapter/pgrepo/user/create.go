package pguser

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

func (us *userRepository) Create(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	const op = "adapter.pgrepo.user.Create"

	const query = `INSERT INTO %s (id, email, login, password_hash, user_role, verify_token)
	VALUES (:id, :email, :login, :password_hash, :user_role, :verify_token)`

	stmt, err := us.db.PrepareNamed(us.table(query))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: failed to prepare statement %w", op, err)
	}
	defer stmt.Close()

	user.ID = uuid.New()
	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%s: failed create user %w", op, err)
	}

	return user.ID, nil
}
