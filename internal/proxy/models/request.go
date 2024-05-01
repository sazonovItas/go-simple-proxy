package models

type Request struct {
	Method      string              `json:"method"`
	Host        string              `json:"host"`
	Path        string              `json:"path"`
	QueryParams map[string][]string `json:"qeury_params,omitempty"`
	PostParams  map[string][]string `json:"post_params,omitempty"`
	Headers     map[string][]string `json:"headers,omitempty"`
	Cookies     map[string]string   `json:"cookies,omitempty"`
	Body        string              `json:"body,omitempty"`
	Raw         string              `json:"raw,omitempty"`
}
