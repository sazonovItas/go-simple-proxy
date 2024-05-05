package proxyutils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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

func Decode(r *models.Request) (*http.Request, error) {
	var body io.Reader
	if r.Body != "" {
		body = strings.NewReader(r.Body)
	}

	// create http request from the model request
	req, err := http.NewRequest(r.Method, fmt.Sprintf("http://%s%s", r.Host, r.Path), body)
	if err != nil {
		return nil, err
	}

	// add query to request
	query := req.URL.Query()
	for param, values := range r.QueryParams {
		for _, value := range values {
			query.Add(param, value)
		}
	}
	req.URL.RawQuery = query.Encode()

	// add headers to request
	for header, values := range r.Headers {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	// add cookies to request
	for cookie, value := range r.Cookies {
		req.AddCookie(&http.Cookie{Name: cookie, Value: value})
	}

	// add post params (form params) to request
	for param, values := range r.PostParams {
		for _, value := range values {
			req.PostForm.Add(param, value)
		}
	}

	return req, nil
}
