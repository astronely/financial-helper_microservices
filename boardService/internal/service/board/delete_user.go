package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) DeleteUser(ctx context.Context, userID int64) error {
	boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board id",
			"error", err.Error(),
		)
		return err
	}

	logger.Debug("DeleteUser",
		"userID", userID,
		"boardID", boardID,
	)
	err = s.boardRepository.DeleteUser(ctx, boardID, userID)
	if err != nil {
		logger.Error("error deleting board:user",
			"error", err.Error(),
		)
		return err
	}

	return nil
}
