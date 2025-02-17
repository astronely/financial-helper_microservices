package access

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/utils"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
	_ "github.com/astronely/financial-helper_microservices/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const authPrefix = "Bearer "

//const accessTokenKey = "access_token_key"

func (s *serv) Check(ctx context.Context, endpointAddress string) (bool, error) {
	if !strings.HasSuffix(endpointAddress, "/Check") {
		return true, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	_, err := utils.VerifyToken(accessToken, []byte(s.tokenConfig.AccessTokenKey()))
	if err != nil {
		logger.Error("Verify access token error",
			"token", accessToken,
			"err", err.Error(),
		)

		cookie, ok := md["grpcgateway-cookie"]

		//logger.Debug("Metadata",
		//	"cookie", cookie)

		if !ok || len(cookie) == 0 {
			return false, status.Errorf(codes.Unauthenticated, "refresh token is not provided")
		}

		refreshToken := strings.TrimPrefix(cookie[0], "token=")

		newAccessToken, err := s.authService.GetAccessToken(ctx, refreshToken)
		if err != nil {
			logger.Error("Error getting access token",
				"err", err.Error(),
			)
			return false, status.Errorf(codes.Unauthenticated, "access tokens is invalid")
		}

		_, err = utils.VerifyToken(newAccessToken, []byte(s.tokenConfig.AccessTokenKey()))
		if err != nil {
			return false, status.Errorf(codes.Unauthenticated, "access tokens is invalid")
		}
	}

	//logger.Debug("Check function in Service Layer",
	//	"Endpoint to check: ", endpointAddress,
	//	"id to check: ", claims.ID,
	//)

	return true, nil
}
