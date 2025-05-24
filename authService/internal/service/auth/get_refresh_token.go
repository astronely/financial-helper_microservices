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

func (s *serv) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.tokenConfig.RefreshTokenKey()))
	if err != nil {
		return "", err
	}

	//logger.Debug("after verifying token")

	token, err := utils.GenerateToken(
		claims.ID,
		model.UserInfo{
			Name:  claims.Username,
			Email: claims.Email,
		},
		[]byte(s.tokenConfig.RefreshTokenKey()), s.tokenConfig.RefreshTokenExpirationTime())

	if err != nil {
		logger.Error("Error in GetRefreshToken",
			"err: ", err.Error(),
		)
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	md := metadata.Pairs("set-cookie", "refreshToken="+token+"; HttpOnly; Path=/; Secure; SameSite=Strict")

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		return "", err
	}

	return token, nil
}
