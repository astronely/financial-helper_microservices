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

	//boardIDStr := md.Get(boardTokenName)[0]
	//logger.Debug("Board ID from context",
	//	"boardID", boardIDStr)

	cookies := md.Get(cookieName)
	if len(cookies) == 0 {
		return -1, errors.New("context doesn't contain cookie")
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
		return -1, errors.New("context doesn't contain token")
	}

	tokenInfo, err := ExtractBoardClaims(tokenString, []byte(key))
	if err != nil {
		logger.Error("ExtractBoardClaims failed",
			"error", err.Error(),
		)
		return -1, err
	}

	return tokenInfo.ID, nil
	//boardID, err := strconv.Atoi(boardIDStr)
	//logger.Debug("Board ID (int) from context",
	//	"boardID", boardID)
	//if err != nil {
	//	return -1, err
	//}
	//return int64(boardID), nil
}
