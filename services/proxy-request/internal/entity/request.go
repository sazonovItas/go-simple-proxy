package entity

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID        uuid.UUID `db:"id"         json:"id,omitempty"`
	UserID    uuid.UUID `db:"user_id"    json:"user_id"`
	ProxyID   uuid.UUID `db:"proxy_id"   json:"proxy_id"`
	RemoteIP  string    `db:"remote_ip"  json:"remote_ip"`
	Host      string    `db:"host"       json:"host"`
	Upload    int64     `db:"upload"     json:"upload"`
	Download  int64     `db:"download"   json:"download"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
