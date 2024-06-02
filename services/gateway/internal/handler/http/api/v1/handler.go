package v1

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

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

func (h *handler) InitRoutes(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initUserRoutes(v1)
		h.initProxyRoutes(v1)
	}
}
