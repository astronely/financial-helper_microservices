package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error) {
	err := s.CheckUserInBoardWithContext(ctx, boardId)
	if err != nil {
		logger.Error("error checking user in board | Service | GetUsers",
			"error", err.Error(),
		)
		return nil, err
	}

	users, err := s.boardRepository.GetUsers(ctx, boardId)
	if err != nil {
		logger.Error("get users error | Service",
			"error", err.Error(),
		)
		return nil, err
	}

	return users, nil
}
