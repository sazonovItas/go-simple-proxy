package grpcuser

type grpcUserRepository interface{}

type UserRepository struct {
	grpcUserRepo grpcUserRepository
}

func New(userRepo grpcUserRepository) *UserRepository {
	return &UserRepository{
		grpcUserRepo: userRepo,
	}
}
