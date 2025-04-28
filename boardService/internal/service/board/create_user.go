package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) CreateUser(ctx context.Context, info *model.BoardUserCreate) (int64, error) {
	userId, err := s.boardRepository.CreateUser(ctx, info)
	if err != nil {
		logger.Error("create user error | Service",
			"error", err.Error(),
		)
		return 0, err
	}

	return userId, nil
}
