package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UpdateUserInfo) (int64, error) {
	id, err := s.userRepository.Update(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, err
}
