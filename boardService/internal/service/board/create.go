package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
)

func (s *serv) Create(ctx context.Context, info *model.BoardCreate) (int64, error) {
	var boardId int64
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return 0, err
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		boardId, errTx = s.boardRepository.Create(ctx, &model.BoardInfo{
			Name:        info.Name,
			Description: info.Description,
			OwnerID:     userID,
		})
		if errTx != nil {
			return errTx
		}

		_, errTx = s.boardRepository.CreateUser(ctx, &model.BoardUserCreate{
			BoardID: boardId,
			UserID:  userID,
			Role:    "admin",
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		logger.Error("create board user error | Service",
			"error", err.Error(),
		)
		return 0, err
	}

	return boardId, nil
}
