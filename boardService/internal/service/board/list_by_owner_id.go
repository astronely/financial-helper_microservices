package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error) {
	return nil, nil
}
