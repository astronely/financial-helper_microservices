package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) CompareUserAndBoard(ctx context.Context) (bool, error) {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("GetUserIdFromContext error",
			"error", err.Error(),
		)
		return false, err
	}
	boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("GetBoardIdFromContext error",
			"error", err.Error(),
		)
		return false, err
	}

	boardUsers, err := s.boardRepository.GetUsers(ctx, boardID)
	if err != nil {
		logger.Error("compare user and board info error",
			"error", err.Error(),
		)
		return false, err
	}

	logger.Debug("Compare | Service",
		"userID", userID,
		"boardID", boardID,
		"boardUsers", boardUsers)

	for _, user := range boardUsers {
		if user.UserID == userID {
			return true, nil
		}
	}

	return false, nil
}
