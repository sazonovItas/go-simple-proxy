package models

import "time"

type Request struct {
	Raw        string    `json:"raw,omitempty" bson:"raw"`
	ReceivedAt time.Time `json:"received_at"   bson:"received_at"`

	Method      string              `json:"method"                 bson:"method"`
	Host        string              `json:"host"                   bson:"host"`
	Path        string              `json:"path"                   bson:"path"`
	QueryParams map[string][]string `json:"qeury_params,omitempty" bson:"query_params"`
	PostParams  map[string][]string `json:"post_params,omitempty"  bson:"post_params"`
	Headers     map[string][]string `json:"headers,omitempty"      bson:"headers"`
	Cookies     map[string]string   `json:"cookies,omitempty"      bson:"cookies"`
	Body        string              `json:"body,omitempty"         bson:"body"`
}

type Response struct {
	Raw        string    `json:"raw,omitempty" bson:"raw"`
	ReceivedAt time.Time `json:"received_at"   bson:"received_at"`

	StatusCode int                 `json:"status_code"       bson:"status_code"`
	Headers    map[string][]string `json:"headers,omitempty" bson:"headers"`
	Cookies    map[string]string   `json:"cookies,omitempty" bson:"cookies"`
	Body       string              `json:"body,omitempty"    bson:"body"`
}

type RequestResponse struct {
	ProcessedAt time.Time `json:"processed_at" bson:"processed_at"`
	Request     Request   `json:"request"      bson:"request"`
	Response    Response  `json:"response"     bson:"response"`
}
