package pgrequest

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *requestRepository) Save(ctx context.Context, r *entity.Request) error {
	const op = "adapter.pgrepo.request.Save"

	const query = `INSERT INTO %s 
		(id, user_id, proxy_id, remote_ip, host, upload, download, created_at) 
	VALUES 
		(:id, :user_id, :proxy_id, :remote_ip, :host, :upload, :download, :created_at)`

	stmt, err := rr.db.PrepareNamedContext(ctx, rr.table(query))
	if err != nil {
		return fmt.Errorf("%s: failed prepare statement: %w", op, err)
	}
	defer stmt.Close()

	r.ID = uuid.New()
	_, err = stmt.ExecContext(ctx, r)
	if err != nil {
		return fmt.Errorf("%s: failed to save request: %w", op, err)
	}

	return nil
}
