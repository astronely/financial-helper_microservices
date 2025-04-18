package access

import (
	"github.com/astronely/financial-helper_microservices/authService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/authService/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
