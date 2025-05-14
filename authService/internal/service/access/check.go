package access

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	_ "github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/authService/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	authPrefix       = "Bearer "
	accessTokenName  = "token"
	refreshTokenName = "refreshToken"
)

//const accessTokenKey = "access_token_key"

func (s *serv) Check(ctx context.Context, endpointAddress string) (bool, error) {
	if strings.HasSuffix(endpointAddress, "login") {
		return true, nil
	}
	if strings.HasSuffix(endpointAddress, "create") {
		return true, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	//logger.Debug("metadata in check",
	//	"metadata", md,
	//)
	if !ok {
		return false, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	logger.Debug("metadata in check",
		"md", md,
	)
	var accessToken, refreshToken string

	cookies := md.Get("grpcgateway-cookie")

	for _, header := range cookies {
		for _, kv := range strings.Split(header, "; ") {
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) != 2 {
				continue
			}
			name, value := parts[0], parts[1]
			if name == accessTokenName {
				accessToken = value
			}
			if name == refreshTokenName {
				refreshToken = value
			}
		}
	}

	if len(accessToken) == 0 || len(refreshToken) == 0 {
		accessToken = md.Get(accessTokenName)[0]
		refreshToken = md.Get(refreshTokenName)[0]
	}

	//if len(accessToken) == 0 {
	//	return false, status.Errorf(codes.Unauthenticated, "cookie is not provided")
	//}

	//var accessToken, refreshToken string
	//for _, c := range cookie {
	//	if c == accessTokenName {
	//		accessToken = c
	//	}
	//	if c == refreshTokenName {
	//		refreshToken = c
	//	}
	//}

	if len(accessToken) == 0 {
		return false, status.Errorf(codes.Unauthenticated, "access token is not provided")
	}

	//authHeader, ok := md["authorization"]
	//if !ok || len(authHeader) == 0 {
	//	return false, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	//}
	//
	//logger.Debug("Authorization token",
	//	"token", authHeader[0],
	//)
	//
	//if !strings.HasPrefix(authHeader[0], authPrefix) {
	//	return false, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	//}
	//
	//accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	//claims :=
	_, err := utils.VerifyToken(accessToken, []byte(s.tokenConfig.AccessTokenKey()))

	if err != nil {
		logger.Error("Verify access token error",
			"token", accessToken,
			"err", err.Error(),
		)
		//refreshToken := md.Get(refreshTokenName)
		if len(refreshToken) == 0 {
			logger.Error("Refresh token is not provided")
			return false, status.Errorf(codes.Unauthenticated, "refresh token is not provided")
		}

		//cookie, ok := md["grpcgateway-cookie"]

		//logger.Debug("Metadata",
		//	"cookie", cookie)

		//if !ok || len(cookie) == 0 {
		//	return false, status.Errorf(codes.Unauthenticated, "refresh token is not provided")
		//}

		//refreshToken := strings.TrimPrefix(cookie[0], "token=")

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
