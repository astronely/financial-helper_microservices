package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Board, error) {
	err := s.CheckUserInBoardWithContext(ctx, id)
	if err != nil {
		logger.Error("error checking user in board | Service | Get",
			"error", err.Error(),
		)
		return nil, err
	}

	board, err := s.boardRepository.Get(ctx, id)
	if err != nil {
		logger.Error("error get board | Service",
			"error", err.Error(),
		)
		return nil, err
	}

	return board, nil
}
