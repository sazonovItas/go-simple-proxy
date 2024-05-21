package entity

import "time"

type Request struct {
	ID            string    `db:"id"              json:"id,omitempty"`
	ProxyID       string    `db:"proxy_id"        json:"proxy_id"`
	ProxyName     string    `db:"proxy_name"      json:"proxy_name"`
	ProxyUserID   string    `db:"proxy_user_id"   json:"proxy_user_id"`
	ProxyUserIP   string    `db:"proxy_user_ip"   json:"proxy_user_ip"`
	ProxyUserName string    `db:"proxy_user_name" json:"proxy_user_name"`
	Host          string    `db:"host"            json:"host"`
	Upload        int64     `db:"upload"          json:"upload"`
	Download      int64     `db:"download"        json:"download"`
	CreatedAt     time.Time `db:"created_at"      json:"created_at"`
}
