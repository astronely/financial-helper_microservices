package auth

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/authService/internal/model"
	"github.com/astronely/financial-helper_microservices/authService/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.tokenConfig.RefreshTokenKey()))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(
		claims.ID,
		model.UserInfo{
			Name:  claims.Username,
			Email: claims.Email,
		},
		[]byte(s.tokenConfig.AccessTokenKey()), s.tokenConfig.AccessTokenExpirationTime())
	if err != nil {
		logger.Error("Error in GetAccessToken",
			"err", err.Error(),
		)
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	err = grpc.SetHeader(ctx, metadata.Pairs("Authorization", "Bearer "+token))
	if err != nil {
		return "", err
	}

	_, err = s.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	return token, nil
}
