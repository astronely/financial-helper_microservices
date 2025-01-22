package user

import (
	"github.com/astronely/financial-helper_microservices/internal/service"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
