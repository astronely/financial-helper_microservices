package user

import (
	"github.com/astronely/financial-helper_microservices/internal/repository"
	def "github.com/astronely/financial-helper_microservices/internal/service"
	"github.com/astronely/financial-helper_microservices/pkg/client/db"
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

func NewMockService(deps ...interface{}) def.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		}
	}

	return &srv
}
