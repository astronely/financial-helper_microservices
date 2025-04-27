package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) ListByUserId(ctx context.Context, userId int64) ([]*model.Board, error) {
	return nil, nil
}
