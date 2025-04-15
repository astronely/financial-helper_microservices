package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Update(ctx context.Context, walletInfo *model.WalletUpdateInfo) (int64, error) {
	id, err := s.walletRepository.Update(ctx, walletInfo)
	if err != nil {
		logger.Error("Failed to update wallet",
			"error", err.Error())
		return -1, err
	}
	return id, nil
}
