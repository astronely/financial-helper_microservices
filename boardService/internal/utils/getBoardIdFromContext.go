package utils

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	boardTokenName = "boardToken"
)

func GetBoardIdFromContext(ctx context.Context, key string) (int64, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return -1, errors.New("context doesn't contain metadata")
	}

	var tokenString string

	cookies := md.Get(cookieName)
	if len(cookies) != 0 {
		for _, header := range cookies {
			for _, kv := range strings.Split(header, "; ") {
				parts := strings.SplitN(kv, "=", 2)
				if len(parts) != 2 {
					continue
				}
				name, value := parts[0], parts[1]
				if name == boardTokenName {
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
	} else {
		token := md.Get(boardTokenName)
		if len(token) == 0 {
			return -1, errors.New("context doesn't contain token")
		}
		if len(token[0]) == 0 {
			return -1, errors.New("context doesn't contain token")
		}
		tokenString = token[0]
	}

	tokenInfo, err := ExtractBoardClaims(tokenString, []byte(key))
	if err != nil {
		logger.Error("ExtractBoardClaims failed",
			"error", err.Error(),
		)
		return -1, err
	}

	return tokenInfo.ID, nil
}
