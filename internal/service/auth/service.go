package auth

import (
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/service"
	"time"
)

const refreshTokenKey = "refresh_token_key"
const accessTokenKey = "access_token_key"
const refreshTokenExpiration = 5 * time.Minute
const accessTimeExpiration = 2 * time.Minute

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{
		authRepository: authRepository,
	}
}
