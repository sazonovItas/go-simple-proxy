package pguser

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
)

func (us *userRepository) UpdatePasswordByResetToken(
	ctx context.Context,
	passwordHash, resetToken string,
) error {
	const op = "adapter.pgrepo.user.NewResetToken"

	const query = "UPDATE %s SET password_hash=$1, reset_token=$2 WHERE reset_token=$2"

	token := sql.NullString{
		Valid:  true,
		String: resetToken,
	}

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statment: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, passwordHash, token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, adapter.ErrResetTokenNotFound)
	}

	return nil
}

func (us *userRepository) NewResetToken(ctx context.Context, email, resetToken string) error {
	const op = "adapter.pgrepo.user.NewResetToken"

	const query = "UPDATE %s SET reset_token=$1 WHERE email=$2"

	token := sql.NullString{
		Valid:  true,
		String: resetToken,
	}

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statment: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, email, token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, adapter.ErrUserNotFound)
	}

	return nil
}

func (us *userRepository) VerifyEmail(ctx context.Context, verifyToken string) error {
	const op = "adapter.pgrepo.user.VerifyEmail"

	const query = "UPDATE %s SET verified=TRUE, verify_token=NULL WHERE verify_token=$1"

	token := sql.NullString{
		Valid:  true,
		String: verifyToken,
	}

	stmt, err := us.db.PreparexContext(ctx, us.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statment: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows == 0 {
		return fmt.Errorf("%s: %w", op, adapter.ErrVerifyTokenNotFound)
	}

	return nil
}
