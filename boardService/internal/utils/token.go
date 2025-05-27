package utils

import (
	"fmt"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(boardID, ownerID int64, secretKey []byte, duration time.Duration) (string, error) {
	claims := model.BoardClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		ID:      boardID,
		OwnerID: ownerID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*model.BoardClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.BoardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.BoardClaims)

	//logger.Debug("VerifyToken",
	//	"secretKey", secretKey,
	//	"claims", claims,
	//	"token", token,
	//)

	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func ExtractUserClaims(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		},
		jwt.WithoutClaimsValidation(),
	)

	if err != nil {
		return nil, err
	}

	logger.Debug("key",
		"secret", secretKey)

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func ExtractBoardClaims(tokenStr string, secretKey []byte) (*model.BoardClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.BoardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		},
		jwt.WithoutClaimsValidation(),
	)

	if err != nil {
		return nil, err
	}

	logger.Debug("key",
		"secret", secretKey)

	claims, ok := token.Claims.(*model.BoardClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
