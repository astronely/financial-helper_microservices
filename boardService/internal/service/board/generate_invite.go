package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

func (s *serv) GenerateInvite(ctx context.Context, info *model.GenerateInviteInfo) (string, error) {
	token, err := s.boardRedisRepository.GenerateInvite(ctx, info)
	if err != nil {
		logger.Error("error generate invite | Board Service",
			"error", err.Error(),
		)
		return "", err
	}
	url := "invite?token=" + token
	return url, nil
}
