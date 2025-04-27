package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error) {
	information, err := s.boardRedisRepository.JoinBoard(ctx, info)
	if err != nil {
		return nil, err
	}

	return information, nil
}
