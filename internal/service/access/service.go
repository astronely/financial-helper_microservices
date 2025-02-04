package access

import "github.com/astronely/financial-helper_microservices/internal/service"

type serv struct {
	authService service.AuthService
}

func NewService(authService service.AuthService) service.AccessService {
	return &serv{
		authService: authService,
	}
}
