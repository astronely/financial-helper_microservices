package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error) {
	id, token, err := s.userRepository.Create(ctx, info, password)
	if err != nil {
		return 0, "", err
	}
	return id, token, nil
}
