package request

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *RequestRepository) GetByID(ctx context.Context, id string) (entity.Request, error) {
	const op = "internal.adapter.pgrepo.request.GetByID"

	const query = "SELECT * FROM %s WHERE id=$1"

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return entity.Request{}, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var request entity.Request
	err = stmt.GetContext(ctx, &request, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return entity.Request{}, ErrRequestNotFound
		default:
			return entity.Request{}, fmt.Errorf("%s: failed exec statement: %w", op, err)
		}
	}

	return entity.Request{}, nil
}
