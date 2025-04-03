package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/utils"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error) {
	id, err := s.userRepository.Create(ctx, info, password)
	if err != nil {
		return 0, "", err
	}

	token, err := utils.GenerateToken(
		id,
		model.UserInfo{
			Name:  info.Name,
			Email: info.Email,
		},
		[]byte(s.tokenConfig.RefreshTokenKey()), s.tokenConfig.RefreshTokenExpirationTime(),
	)

	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}
