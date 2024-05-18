package entity

import "time"

type Request struct {
	ID        string    `json:"id"`
	Host      string    `json:"host"`
	Upload    int64     `json:"upload"`
	Download  int64     `json:"download"`
	CreatedAt time.Time `json:"created_at"`
}
