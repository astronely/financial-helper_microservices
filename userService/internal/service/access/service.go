package access

import (
	"github.com/astronely/financial-helper_microservices/userService/internal/config"
	"github.com/astronely/financial-helper_microservices/userService/internal/service"
)

type serv struct {
	authService service.AuthService
	tokenConfig config.TokenConfig
}

func NewService(authService service.AuthService, tokenConfig config.TokenConfig) service.AccessService {
	return &serv{
		authService: authService,
		tokenConfig: tokenConfig,
	}
}
