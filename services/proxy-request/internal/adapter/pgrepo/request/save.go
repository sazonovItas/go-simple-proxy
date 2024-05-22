package pgrequest

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *RequestRepository) Save(ctx context.Context, request *entity.Request) error {
	const op = "internal.adapter.pgrepo.request.Save"

	const query = `INSERT INTO %s 
		(id, proxy_id, proxy_name, proxy_user_id, proxy_user_ip, proxy_user_name, host, upload, download, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	stmt, err := rr.db.PreparexContext(ctx, rr.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	id := uuid.New()
	_, err = stmt.ExecContext(
		ctx,
		id,
		request.ProxyID,
		request.ProxyName,
		request.ProxyUserID,
		request.ProxyUserIP,
		request.ProxyUserName,
		request.Host,
		request.Upload,
		request.Download,
		request.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s: failed to exec statement: %w", op, err)
	}
	request.ID = id.String()

	return nil
}
