package entity

import "time"

type Request struct {
	ID            string    `json:"id"`
	ProxyID       string    `json:"proxy_id"`
	ProxyName     string    `json:"proxy_name"`
	ProxyUserID   string    `json:"proxy_user_id"`
	ProxyUserIP   string    `json:"proxy_user_ip"`
	ProxyUserName string    `json:"proxy_user_name"`
	Host          string    `json:"host"`
	Upload        int64     `json:"upload"`
	Download      int64     `json:"download"`
	CreatedAt     time.Time `json:"created_at"`
}
