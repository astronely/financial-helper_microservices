package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error) {
	return nil, nil
}
