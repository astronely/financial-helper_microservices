package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) GenerateInvite(ctx context.Context) (string, error) {
	boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board id",
			"error", err.Error(),
		)
		return "", err
	}

	err = s.CheckUserInBoardWithContext(ctx, boardID)
	if err != nil {
		logger.Error("error checking user in board | Service | GenerateInvite",
			"error", err.Error(),
		)
		return "", err
	}

	info := converter.ToGenerateInviteInfo(boardID, "editor")
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
