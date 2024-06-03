package v1

import (
	"time"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

type registerResponse struct{}

func newRegisterResponse() *registerResponse {
	return &registerResponse{}
}

type loginResponse struct {
	Token string `json:"token"`
}

func newLoginResponse(token string) *loginResponse {
	return &loginResponse{
		Token: token,
	}
}

type userResponse struct {
	User struct {
		Email     string    `json:"email"`
		Login     string    `json:"login"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"user"`
}

func newUserResponse(u *entity.User) *userResponse {
	r := new(userResponse)
	r.User.Email = u.Email
	r.User.Login = u.Login
	r.User.CreatedAt = u.CreatedAt
	return r
}

type requestsResponse struct {
	Requests []*entity.Request `json:"requests"`
}

func newUserProxyRequestsReqsponse(requests []*entity.Request) *requestsResponse {
	return &requestsResponse{
		Requests: requests,
	}
}

type proxyInfoResponse struct {
	Proxies []*entity.Proxy `json:"proxies"`
}

func newProxyInfoReqsponse(proxies []*entity.Proxy) *proxyInfoResponse {
	return &proxyInfoResponse{
		Proxies: proxies,
	}
}
