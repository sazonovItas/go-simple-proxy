package proxysvc

import (
	"context"
	"fmt"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/lib/hasher"
)

func (ps *ProxyService) Login(
	ctx context.Context,
	login, password string,
) (entity.Token, error) {
	const op = "service.proxy.Login"

	tokenKey := hasher.Hash(login + ":" + password)
	token, err := ps.tokenRepo.Get(tokenKey)
	if err == nil {

		_, err := ps.authRepo.Validate(ctx, token.TokenString)
		if err == nil {
			return token, nil
		}

		ps.log.Warn("failed validate token", slogger.Err(err))
	}

	tokenString, err := ps.authRepo.Login(ctx, login, password)
	if err != nil {
		ps.log.Error("failed to login user", slogger.Err(err))

		return entity.Token{}, fmt.Errorf("%s: %w", op, err)
	}

	id, err := ps.authRepo.Validate(ctx, tokenString)
	if err != nil {
		ps.log.Error("failed validate token", slogger.Err(err))

		return entity.Token{}, fmt.Errorf("%s: %w", op, err)
	}

	token = entity.Token{UserID: id, TokenString: tokenString}
	ps.tokenRepo.Set(tokenKey, token, 0)

	return token, nil
}
