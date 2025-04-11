package service

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

type WalletService interface {
	Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Wallet, error)
	List(ctx context.Context, limit, offset uint64) ([]*model.Wallet, error)
	Update(ctx context.Context, walletInfo *model.UpdateWalletInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type TransactionService interface{}
