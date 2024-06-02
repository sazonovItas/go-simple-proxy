package entity

import "time"

type User struct {
	Email     string    `json:"email,omitempty"`
	Login     string    `json:"login,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
