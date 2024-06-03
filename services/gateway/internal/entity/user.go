package entity

import "time"

type User struct {
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Passsword string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
