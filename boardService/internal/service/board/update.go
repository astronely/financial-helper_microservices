package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.BoardUpdate) (int64, error) {
	boardId, err := s.boardRepository.Update(ctx, info)
	if err != nil {
		logger.Error("error update board | Service",
			"error", err.Error(),
		)
		return 0, err
	}

	return boardId, nil
}
