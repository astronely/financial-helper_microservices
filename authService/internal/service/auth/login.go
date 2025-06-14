package auth

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/authService/internal/model"
	"github.com/astronely/financial-helper_microservices/authService/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *serv) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	//logger.Debug("User in service layer, login",
	//	"user", user,
	//)

	if !utils.VerifyPassword(user.Info.Password, password) {
		return "", errors.New("invalid password")
	}

	refreshToken, err := utils.GenerateToken(
		user.ID,
		model.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
		},
		[]byte(s.tokenConfig.RefreshTokenKey()), s.tokenConfig.RefreshTokenExpirationTime(),
	)

	accessToken, err := utils.GenerateToken(
		user.ID,
		model.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
		},
		[]byte(s.tokenConfig.AccessTokenKey()), s.tokenConfig.AccessTokenExpirationTime(),
	)

	//err = grpc.SetHeader(ctx, metadata.Pairs("Authorization", "Bearer "+accessToken))

	if err != nil {
		return "", err
	}

	mdRefreshToken := metadata.Pairs("set-cookie", "refreshToken="+refreshToken+"; HttpOnly; Path=/;SameSite=Lax;")
	mdAccessToken := metadata.Pairs("set-cookie", "token="+accessToken+"; HttpOnly; Path=/; SameSite=Lax")

	err = grpc.SetHeader(ctx, mdRefreshToken)
	if err != nil {
		return "", err
	}
	err = grpc.SetHeader(ctx, mdAccessToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
