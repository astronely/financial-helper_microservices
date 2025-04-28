package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.boardRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("error deleting board | Service",
			"error", err.Error(),
		)
		return err
	}

	return nil
}
