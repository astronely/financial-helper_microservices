package auth

import (
	"github.com/astronely/financial-helper_microservices/internal/service"
	desc "github.com/astronely/financial-helper_microservices/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
