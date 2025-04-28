package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error) {
	boardInviteInfo, err := s.boardRedisRepository.JoinBoard(ctx, info)
	if err != nil {
		logger.Error("join board error | Redis",
			"error", err.Error(),
		)
		return nil, err
	}

	_, err = s.boardRepository.CreateUser(ctx, &model.BoardUserCreate{
		BoardID: boardInviteInfo.BoardID,
		UserID:  info.ID,
		Role:    boardInviteInfo.Role,
	})

	if err != nil {
		logger.Error("join board | Create user error",
			"error", err.Error(),
		)
		return nil, err
	}

	return boardInviteInfo, nil
}
