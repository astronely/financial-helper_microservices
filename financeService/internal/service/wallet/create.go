package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error) {
	id, err := s.walletRepository.Create(ctx, walletInfo)
	if err != nil {
		logger.Error("Failed to create new wallet",
			"error", err.Error())
		return 0, err
	}
	return id, nil
}
