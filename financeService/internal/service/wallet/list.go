package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) List(ctx context.Context, boardID int64) ([]*model.Wallet, error) {
	wallets, err := s.walletRepository.List(ctx, boardID)
	if err != nil {
		logger.Error("Failed to get list of wallets",
			"error", err.Error())
		return nil, err
	}

	return wallets, nil
}
