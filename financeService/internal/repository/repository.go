package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

type WalletRepository interface {
	Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Wallet, error)
	List(ctx context.Context, limit, offset uint64) ([]*model.Wallet, error)
	Update(ctx context.Context, walletInfo *model.UpdateWalletInfo) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type TransactionRepository interface {
	CreateTransactionDetails(ctx context.Context, transactionDetails *model.TransactionDetailsInfo) (int64, error)
	CreateTransaction(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsId int64) (int64, error)
}
