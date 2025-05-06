package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) ListByUserId(ctx context.Context) ([]*model.Board, error) {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id",
			"error", err.Error(),
		)
		return nil, err
	}

	boards, err := s.boardRepository.ListByUserId(ctx, userID)
	if err != nil {
		logger.Error("error list by userId | Service",
			"error", err.Error(),
		)
		return nil, err
	}

	return boards, nil
}
