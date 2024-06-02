package entity

type Request struct {
	ID       string `json:"id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	ProxyID  string `json:"proxy_id,omitempty"`
	RemoteIP string `json:"remote_ip,omitempty"`
	Host     string `json:"host,omitempty"`
	Upload   int64  `json:"upload,omitempty"`
	Download int64  `json:"download,omitempty"`
}
