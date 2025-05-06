package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
)

func (s *serv) CompareUserAndBoardRaw(ctx context.Context, userId int64, boardId int64) (bool, error) {
	boardUsers, err := s.boardRepository.GetUsers(ctx, boardId)
	if err != nil {
		logger.Error("compare user and board info error",
			"error", err.Error(),
		)
		return false, err
	}

	logger.Debug("Compare | Service",
		"userID", userId,
		"boardID", boardId,
		"boardUsers", boardUsers)

	for _, user := range boardUsers {
		if user.UserID == userId {
			return true, nil
		}
	}

	return false, nil
}
