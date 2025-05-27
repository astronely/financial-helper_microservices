package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.CheckUserInBoardWithContext(ctx, id)
	if err != nil {
		logger.Error("error checking user in board | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	board, err := s.boardRepository.Get(ctx, id)
	if err != nil {
		logger.Error("error getting board | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	if userId != board.Info.OwnerID {
		//logger.Error("error checking user is owner of board | Service | Delete",
		//	"userId", userId,
		//	"ownerId", board.Info.OwnerID)
		err = s.boardRepository.DeleteUser(ctx, id, userId)
		if err != nil {
			logger.Error("error deleting user from board | Service | Delete",
				"error", err.Error(),
			)
			return err
		}
		return nil
		//return errors.New("not allowed")
	}

	err = s.boardRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("error deleting board | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	return nil
}
