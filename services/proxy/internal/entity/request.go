package entity

import "time"

type Request struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id"`
	ProxyID   string    `json:"proxy_id"`
	RemoteIP  string    `json:"remote_ip"`
	Host      string    `json:"host"`
	Upload    int64     `json:"upload"`
	Download  int64     `json:"download"`
	CreatedAt time.Time `json:"created_at"`
}
