package v1

import (
	"context"
	"time"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

type RequestService interface {
	TimestampAndUserId(
		ctx context.Context,
		from, to time.Time,
		userId string,
	) ([]*entity.Request, error)
}

type UserService interface {
	Register(ctx context.Context, email, login, password string) (id string, err error)
	Login(ctx context.Context, login, password string) (token string, err error)
	ValidateToken(ctx context.Context, token string) (tokenClaims *entity.Token, err error)

	User(ctx context.Context, id string) (user *entity.User, err error)
}

type ProxyService interface {
	ProxyInfo(ctx context.Context) (proxies []*entity.Proxy, err error)
}

type handler struct {
	userSvc  UserService
	proxySvc ProxyService
	reqSvc   RequestService
}

func NewHandler(userSvc UserService, reqSvc RequestService, proxySvc ProxyService) *handler {
	return &handler{
		userSvc:  userSvc,
		proxySvc: proxySvc,
		reqSvc:   reqSvc,
	}
}
