package auth

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	token, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())

	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: token,
	}, nil
}
