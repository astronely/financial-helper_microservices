package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/shopspring/decimal"
)

type WalletRepository interface {
	Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Wallet, error)
	List(ctx context.Context, boardID int64) ([]*model.Wallet, error)
	Update(ctx context.Context, walletInfo *model.WalletUpdateInfo) (int64, error)
	UpdateBalance(ctx context.Context, id int64, value decimal.Decimal, txType string) error
	Delete(ctx context.Context, id int64) error
}

type TransactionRepository interface {
	CreateTransactionDetails(ctx context.Context, transactionDetails *model.TransactionDetailsInfo) (int64, error)
	CreateTransaction(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsId int64) (int64, error)
	Get(ctx context.Context, id int64, filters map[string]interface{}) (*model.Transaction, error)
	List(ctx context.Context, limit, offset uint64, filters map[string]interface{}) ([]*model.Transaction, error)
	UpdateInfo(ctx context.Context, updateInfo *model.TransactionInfoUpdate) (*model.TransactionInfo, error)
	UpdateDetails(ctx context.Context, updateInfo *model.TransactionDetailsUpdate) (int64, error)
	Categories(ctx context.Context) ([]*model.TransactionCategory, error)
	Delete(ctx context.Context, id int64) error
}
