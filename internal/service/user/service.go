package user

import (
	"github.com/astronely/financial-helper_microservices/internal/client/db"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	def "github.com/astronely/financial-helper_microservices/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) *serv {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
