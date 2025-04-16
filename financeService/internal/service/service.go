package service

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

type WalletService interface {
	Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Wallet, error)
	List(ctx context.Context, limit, offset uint64) ([]*model.Wallet, error)
	Update(ctx context.Context, walletInfo *model.WalletUpdateInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type TransactionService interface {
	Create(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsInfo *model.TransactionDetailsInfo) (int64, error)
	Get(ctx context.Context, id int64, filters map[string]interface{}) (*model.Transaction, error)
	List(ctx context.Context, limit, offset uint64, filters map[string]interface{}) ([]*model.Transaction, error)
	Update(ctx context.Context,
		updateInfo *model.TransactionInfoUpdate,
		updateDetails *model.TransactionDetailsUpdate) (int64, error)
	Categories(ctx context.Context) ([]*model.TransactionCategory, error)
}
