package utils

import (
	"fmt"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

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

	//logger.Debug("key",
	//	"secret", secretKey)

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

	//logger.Debug("key",
	//	"secret", secretKey)

	claims, ok := token.Claims.(*model.BoardClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
