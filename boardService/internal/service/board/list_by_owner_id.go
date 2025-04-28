package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error) {
	boards, err := s.boardRepository.ListByOwnerId(ctx, ownerId)
	if err != nil {
		logger.Error("error list by owner id | Service",
			"error", err.Error(),
		)
		return nil, err
	}

	return boards, nil
}
