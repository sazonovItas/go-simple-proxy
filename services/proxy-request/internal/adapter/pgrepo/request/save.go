package requestrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *RequestRepository) Save(ctx context.Context, request *entity.Request) error {
	const op = "adapter.pgrepo.request.Save"

	const query = `INSERT INTO %s 
		(id, proxy_id, proxy_name, proxy_user_id, proxy_user_ip, proxy_user_name, host, upload, download, created_at) 
	VALUES 
		(:id, :proxy_id, :proxy_name, :proxy_user_id, :proxy_user_ip, :proxy_user_name, :host, :upload, :download, :created_at)`

	stmt, err := rr.db.PrepareNamedContext(ctx, rr.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	request.ID = uuid.New()
	_, err = stmt.ExecContext(ctx, request)
	if err != nil {
		return fmt.Errorf("%s: failed to exec statement: %w", op, err)
	}

	return nil
}
