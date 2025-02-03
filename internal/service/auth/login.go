package auth

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/utils"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *serv) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	logger.Debug("User in service layer, login",
		"user", user,
	)

	if !utils.VerifyPassword(user.Info.Password, password) {
		return "", errors.New("invalid password")
	}

	refreshToken, err := utils.GenerateToken(
		user.ID,
		model.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
		},
		[]byte(refreshTokenKey), refreshTokenExpiration,
	)

	accessToken, err := utils.GenerateToken(
		user.ID,
		model.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
		},
		[]byte(accessTokenKey), accessTimeExpiration,
	)

	err = grpc.SetHeader(ctx, metadata.Pairs("Authorization", "Bearer "+accessToken))

	if err != nil {
		return "", err
	}

	md := metadata.Pairs("set-cookie", "token="+refreshToken+"; HttpOnly; Path=/; Secure; SameSite=Strict")

	err = grpc.SendHeader(ctx, md)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
