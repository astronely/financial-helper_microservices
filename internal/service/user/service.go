package user

import (
	"github.com/astronely/financial-helper_microservices/internal/repository"
	def "github.com/astronely/financial-helper_microservices/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}
