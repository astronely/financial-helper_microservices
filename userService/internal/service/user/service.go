package user

import (
	"github.com/astronely/financial-helper_microservices/userService/internal/config"
	"github.com/astronely/financial-helper_microservices/userService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/userService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	tokenConfig    config.TokenConfig
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager, tokenConfig config.TokenConfig) def.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
		tokenConfig:    tokenConfig,
	}
}

func NewMockService(deps ...interface{}) def.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		case config.TokenConfig:
			srv.tokenConfig = s
		}
	}

	return &srv
}
