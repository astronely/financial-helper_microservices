package auth

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/authService/pkg/auth_v1"
)

func (i *Implementation) Logout(ctx context.Context, req *desc.LogoutRequest) (*desc.LogoutResponse, error) {
	err := i.authService.Logout(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.LogoutResponse{}, nil
}
