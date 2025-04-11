package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.walletRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("Failed to delete wallet",
			"error", err.Error())
	}
	return err
}
