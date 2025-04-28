package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) ListByUserId(ctx context.Context, userId int64) ([]*model.Board, error) {
	boards, err := s.boardRepository.ListByUserId(ctx, userId)
	if err != nil {
		logger.Error("error list by userId | Service",
			"error", err.Error(),
		)
		return nil, err
	}

	return boards, nil
}
