package common

import (
	"net/http"
	"net/url"

	"github.com/sazonovItas/go-simple-proxy/internal/proxy/models"
)

func ParseRequest(r *http.Request) *models.Request {
	req := &models.Request{
		Host:   r.Host,
		Method: r.Method,
		Path:   r.URL.Path,
	}

	queryParamVals := make(url.Values)
	for k, values := range r.URL.Query() {
		queryParamVals[k] = append(queryParamVals[k], values...)
	}
	req.QueryParams = queryParamVals

	headers := make(http.Header)
	for k, values := range r.Header {
		headers[k] = append(headers[k], values...)
	}
	req.Headers = headers

	cookies := make(map[string]string)
	for _, v := range r.Cookies() {
		cookies[v.Name] = v.Value
	}
	req.Cookies = cookies

	if err := r.ParseForm(); err != nil {
		postParams := make(url.Values)
		for k, values := range r.PostForm {
			postParams[k] = append(postParams[k], values...)
		}
		req.PostParams = postParams
	}

	return req
}
