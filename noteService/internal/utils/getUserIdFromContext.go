package utils

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	userTokenName = "token"
	cookieName    = "grpcgateway-cookie"
)

func GetUserIdFromContext(ctx context.Context, key string) (int64, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return -1, errors.New("context doesn't contain metadata")
	}
	logger.Debug("md",
		"metadata", md,
	)
	cookies := md.Get(cookieName)
	if len(cookies) == 0 {
		return -1, errors.New("context doesn't contain cookie")
	}
	logger.Debug("cookies",
		"cookies", cookies)

	var tokenString string
	for _, header := range cookies {
		for _, kv := range strings.Split(header, "; ") {
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) != 2 {
				continue
			}
			name, value := parts[0], parts[1]
			if name == userTokenName {
				tokenString = value
				break
			}
		}
		if tokenString != "" {
			break
		}
	}

	if tokenString == "" {
		return -1, errors.New("context doesn't contain token")
	}

	logger.Debug("token info",
		"token", tokenString,
		"secret", key)
	tokenInfo, err := ExtractUserClaims(tokenString, []byte(key))
	if err != nil {
		logger.Error("ExtractUserClaims failed",
			"error", err.Error(),
		)
		return -1, err
	}

	return tokenInfo.ID, nil
}
