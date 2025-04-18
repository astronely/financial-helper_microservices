package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) List(ctx context.Context, limit, offset uint64) ([]*model.Wallet, error) {
	wallets, err := s.walletRepository.List(ctx, limit, offset)
	if err != nil {
		logger.Error("Failed to get list of wallets",
			"error", err.Error())
		return nil, err
	}

	return wallets, nil
}
