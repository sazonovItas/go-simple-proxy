package models

type Response struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers,omitempty"`
	Cookies    map[string]string   `json:"cookies,omitempty"`
	Body       string              `json:"body,omitempty"`
	Raw        string              `json:"raw,omitempty"`
}
