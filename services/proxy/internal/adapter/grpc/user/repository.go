package userrepo

import "google.golang.org/grpc"

type rpcUserRepository struct{}

type UserRepository struct {
	rpcUserRepo rpcUserRepository
}

func NewUserRepository(cli *grpc.ClientConn) *UserRepository {
	return &UserRepository{}
}
