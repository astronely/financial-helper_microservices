package board

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) CheckUserInBoardWithContext(ctx context.Context, boardID int64) error {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return err
	}

	result, err := s.CompareUserAndBoardRaw(ctx, userID, boardID)
	if err != nil {
		logger.Error("error compare user and board",
			"error", err.Error(),
		)
		return err
	}

	if !result {
		logger.Error("board doesnt contains this user")
		return errors.New("board doesnt contains this user")
	}

	return nil
}
