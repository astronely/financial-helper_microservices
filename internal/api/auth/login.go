package auth

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/pkg/auth_v1"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	token, err := i.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		logger.Error("Error login",
			"email", req.GetEmail(),
		)
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: token,
	}, nil

}
