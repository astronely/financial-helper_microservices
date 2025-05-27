package utils

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"google.golang.org/grpc/metadata"
	"strings"
)

func GetBoardFromContext(ctx context.Context, key string) (*model.BoardClaims, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return nil, errors.New("context doesn't contain metadata")
	}

	cookies := md.Get(cookieName)
	if len(cookies) == 0 {
		return nil, errors.New("context doesn't contain cookie")
	}

	var tokenString string
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
		return nil, errors.New("context doesn't contain token")
	}

	tokenInfo, err := ExtractBoardClaims(tokenString, []byte(key))
	if err != nil {
		logger.Error("ExtractBoardClaims failed",
			"error", err.Error(),
		)
		return nil, err
	}

	return tokenInfo, nil
}
