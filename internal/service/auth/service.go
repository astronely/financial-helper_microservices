package auth

import (
	"github.com/astronely/financial-helper_microservices/internal/config"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/service"
)

//const refreshTokenKey = "refresh_token_key"
//const accessTokenKey = "access_token_key"
//const refreshTokenExpiration = 5 * time.Minute
//const accessTimeExpiration = 2 * time.Minute

type serv struct {
	authRepository repository.AuthRepository
	tokenConfig    config.TokenConfig
}

func NewService(authRepository repository.AuthRepository, tokenConfig config.TokenConfig) service.AuthService {
	return &serv{
		authRepository: authRepository,
		tokenConfig:    tokenConfig,
	}
}
