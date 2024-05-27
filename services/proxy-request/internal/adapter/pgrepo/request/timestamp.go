package pgrequest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *requestRepository) Timestamp(
	ctx context.Context,
	from, to time.Time,
) ([]entity.Request, error) {
	const op = "adapter.pgrepo.request.Timestamp"

	const query = `SELECT * FROM %s WHERE created_at BETWEEN $1 AND $2
	ORDER BY created_at DESC`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, from, to)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed select requests: %w", op, err)
		}
	}

	return requests, nil
}

func (rr *requestRepository) TimestampAndUserId(
	ctx context.Context,
	from, to time.Time,
	userId uuid.UUID,
) ([]entity.Request, error) {
	const op = "adapter.pgrepo.request.TimestampAndUserId"

	const query = `SELECT * FROM %s WHERE (created_at BETWEEN $1 AND $2) AND user_id=$3
	ORDER BY created_at DESC`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, from, to, userId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed select requests: %w", op, err)
		}
	}

	return requests, nil
}

func (rr *requestRepository) TimestampAndProxyId(
	ctx context.Context,
	from, to time.Time,
	proxyId uuid.UUID,
) ([]entity.Request, error) {
	const op = "adapter.pgrepo.request.TimestampAndProxyId"

	const query = `SELECT * FROM %s WHERE (created_at BETWEEN $1 AND $2) AND proxy_id=$3
	ORDER BY created_at DESC`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, from, to, proxyId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed select requests: %w", op, err)
		}
	}

	return requests, nil
}
