package entity

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID            uuid.UUID `db:"id"              json:"id,omitempty"`
	ProxyID       uuid.UUID `db:"proxy_id"        json:"proxy_id"`
	ProxyName     string    `db:"proxy_name"      json:"proxy_name"`
	ProxyUserID   uuid.UUID `db:"proxy_user_id"   json:"proxy_user_id"`
	ProxyUserIP   string    `db:"proxy_user_ip"   json:"proxy_user_ip"`
	ProxyUserName string    `db:"proxy_user_name" json:"proxy_user_name"`
	Host          string    `db:"host"            json:"host"`
	Upload        int64     `db:"upload"          json:"upload"`
	Download      int64     `db:"download"        json:"download"`
	CreatedAt     time.Time `db:"created_at"      json:"created_at"`
}
