package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	SimpleUser = "user"
	AdminUser  = "admin"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	Login        string    `db:"login"`
	PasswordHash string    `db:"password_hash"`
	UserRole     string    `db:"user_role"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeletedAt    time.Time `db:"deleted_at"`
}

func (u User) IsDeleted() bool {
	return u.DeletedAt.After(u.CreatedAt)
}

func (u User) IsAdmin() bool {
	return u.UserRole == AdminUser
}
