package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

func (s *serv) List(ctx context.Context, limit uint64, offset uint64) ([]*model.User, error) {
	users, err := s.userRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}
