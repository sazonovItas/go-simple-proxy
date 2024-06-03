package entity

type Request struct {
	RemoteIP string `json:"remote_ip,omitempty"`
	Host     string `json:"host,omitempty"`
	Upload   int64  `json:"upload,omitempty"`
	Download int64  `json:"download,omitempty"`
}
