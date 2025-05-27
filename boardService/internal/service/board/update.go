package board

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) Update(ctx context.Context, info *model.BoardUpdate) (int64, error) {
	err := s.CheckUserInBoardWithContext(ctx, info.ID)
	if err != nil {
		logger.Error("error checking user in board | Service | Update",
			"error", err.Error(),
		)
		return 0, err
	}

	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Update",
			"error", err.Error(),
		)
		return 0, err
	}

	board, err := s.boardRepository.Get(ctx, info.ID)
	if err != nil {
		logger.Error("error getting board | Service | Update",
			"error", err.Error(),
		)
		return 0, err
	}

	if userId != board.Info.OwnerID {
		logger.Error("error checking user is owner of board | Service | Update",
			"userId", userId,
			"ownerId", board.Info.OwnerID,
		)
		return 0, errors.New("not allowed")
	}

	boardId, err := s.boardRepository.Update(ctx, info)
	if err != nil {
		logger.Error("error update board | Service",
			"error", err.Error(),
		)
		return 0, err
	}

	return boardId, nil
}
