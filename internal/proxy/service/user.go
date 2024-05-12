package service

import (
	"context"

	"github.com/sazonovItas/go-simple-proxy/internal/proxy/entity"
)

type userRepository interface {
	FindByUsername(ctx context.Context, username string) (entity.User, error)
}

type UserService struct {
	userRepo userRepository
}

func NewUserService(userRepository userRepository) *UserService {
	return &UserService{
		userRepo: userRepository,
	}
}
