package requestrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *RequestRepository) Timestamp(
	ctx context.Context,
	timestamp time.Time,
	limit int,
) ([]entity.Request, error) {
	const op = "internal.adapter.pgrepo.request.GetByTimestamp"

	const query = `SELECT * FROM %s WHERE created_at<$1 
	ORDER BY created_at DESC LIMIT $2`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, timestamp, limit)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed exec statement: %w", op, err)
		}
	}

	return requests, nil
}

func (rr *RequestRepository) GetByProxyIDAndTimestamp(
	ctx context.Context,
	timestamp time.Time,
	proxyId string,
	limit int,
) ([]entity.Request, error) {
	const op = "internal.adapter.pgrepo.request.GetByProxyUserIDAndTimestamp"

	const query = `SELECT * FROM %s WHERE created_at<$1 AND proxy_id=$2 
	ORDER BY created_at DESC LIMIT $3`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, timestamp, proxyId, limit)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed exec statement: %w", op, err)
		}
	}

	return requests, nil
}

func (rr *RequestRepository) GetByProxyUserIDAndTimestamp(
	ctx context.Context,
	timestamp time.Time,
	proxyUserId string,
	limit int,
) ([]entity.Request, error) {
	const op = "internal.adapter.pgrepo.request.GetByProxyUserIDAndTimestamp"

	const query = `SELECT * FROM %s WHERE created_at<$1 AND proxy_user_id=$2 
	ORDER BY created_at DESC LIMIT $3`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, timestamp, proxyUserId, limit)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed exec statement: %w", op, err)
		}
	}

	return requests, nil
}

func (rr *RequestRepository) GetByHostAndTimestamp(
	ctx context.Context,
	timestamp time.Time,
	host string,
	limit int,
) ([]entity.Request, error) {
	const op = "internal.adapter.pgrepo.request.GetByHostAndTimestamp"

	const query = `SELECT * FROM %s WHERE created_at<$1 AND host=$2 
	ORDER BY created_at DESC LIMIT $3`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return nil, fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var requests []entity.Request
	err = stmt.SelectContext(ctx, &requests, timestamp, host, limit)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, adapter.ErrRequestNotFound
		default:
			return nil, fmt.Errorf("%s: failed exec statement: %w", op, err)
		}
	}

	return requests, nil
}
