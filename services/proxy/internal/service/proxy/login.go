package proxysvc

import (
	"context"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (ps *ProxyService) Login(
	ctx context.Context,
	username, passwordHash string,
) (*entity.User, error) {
	return &entity.User{
		ID:           uuid.NewString(),
		Username:     username,
		PasswordHash: passwordHash,
	}, nil
}
