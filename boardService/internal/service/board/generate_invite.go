package board

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) GenerateInvite(ctx context.Context) (string, error) {
	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Update",
			"error", err.Error(),
		)
		return "", err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context | Service | Update",
			"error", err.Error(),
		)
		return "", err
	}

	if userId != board.OwnerID {
		return "", errors.New("not allowed")
	}

	err = s.CheckUserInBoardWithContext(ctx, board.ID)
	if err != nil {
		logger.Error("error checking user in board | Service | GenerateInvite",
			"error", err.Error(),
		)
		return "", err
	}

	info := converter.ToGenerateInviteInfo(board.ID, "editor")
	token, err := s.boardRedisRepository.GenerateInvite(ctx, info)
	if err != nil {
		logger.Error("error generate invite | Board Service",
			"error", err.Error(),
		)
		return "", err
	}
	//url := "invite?token=" + token
	return token, nil
}
