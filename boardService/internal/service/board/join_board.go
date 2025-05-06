package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error) {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return nil, err
	}

	boardInviteInfo, err := s.boardRedisRepository.JoinBoard(ctx, info)
	if err != nil {
		logger.Error("join board error | Redis",
			"error", err.Error(),
		)
		return nil, err
	}

	_, err = s.boardRepository.CreateUser(ctx, &model.BoardUserCreate{
		BoardID: boardInviteInfo.BoardID,
		UserID:  userID,
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
