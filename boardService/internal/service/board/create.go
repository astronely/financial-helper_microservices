package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.BoardInfo) (int64, error) {
	var boardId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		boardId, errTx = s.boardRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.boardRepository.CreateUser(ctx, &model.BoardUserCreate{
			BoardID: boardId,
			UserID:  info.OwnerID,
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
