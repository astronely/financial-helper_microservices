package auth

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	accessTokenName  = "token"
	refreshTokenName = "refreshToken"
	boardTokenName   = "boardToken"
)

func (s *serv) Logout(ctx context.Context) error {
	mdRefreshToken := metadata.Pairs("set-cookie", refreshTokenName+"=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; HttpOnly; Path=/; Secure=false; SameSite=None;")
	err := grpc.SetHeader(ctx, mdRefreshToken)
	if err != nil {
		logger.Error("error set expired refresh token",
			"error", err.Error(),
		)
		return err
	}

	mdAccessToken := metadata.Pairs("set-cookie", accessTokenName+"=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; HttpOnly; Path=/; Secure=false; SameSite=None")
	err = grpc.SetHeader(ctx, mdAccessToken)
	if err != nil {
		logger.Error("error set expired access token")
		return err
	}

	mdBoardToken := metadata.Pairs("set-cookie", boardTokenName+"=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; HttpOnly; Path=/; Secure=false; SameSite=None;")
	err = grpc.SetHeader(ctx, mdBoardToken)
	if err != nil {
		logger.Error("error set expired board token",
			"error", err.Error(),
		)
		return err
	}

	return nil
}
