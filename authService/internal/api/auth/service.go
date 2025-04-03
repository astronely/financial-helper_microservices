package auth

import (
	"github.com/astronely/financial-helper_microservices/authService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/authService/pkg/auth_v1"
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
