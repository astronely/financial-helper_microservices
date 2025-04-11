package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Update(ctx context.Context, walletInfo *model.UpdateWalletInfo) (int64, error) {
	id, err := s.walletRepository.Update(ctx, walletInfo)
	if err != nil {
		return -1, err
	}
	return id, nil
}
