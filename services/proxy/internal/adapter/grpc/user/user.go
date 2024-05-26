package grpcuser

import (
	"context"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (ur *UserRepository) UserByName(ctx context.Context, name string) (*entity.User, error) {
	panic("method is not impelemented")
}
