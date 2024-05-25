package proxysvc

import (
	"context"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (ps *ProxyService) Save(ctx context.Context, r *entity.Request) error {
	return ps.requestRepo.Save(ctx, r)
}
