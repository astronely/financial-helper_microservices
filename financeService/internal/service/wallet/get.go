package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Wallet, error) {
	wallet, err := s.walletRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
