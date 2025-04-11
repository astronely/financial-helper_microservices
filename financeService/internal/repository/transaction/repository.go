package transaction

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

const (
	transactionTableName        = "transactions"
	transactionDetailsTableName = "transaction_details"

	idColumn        = "id"
	ownerIdColumn   = "owner_id"
	WalletIdColumn  = "wallet_id"
	boardIdColumn   = "board_id"
	sumColumn       = "sum"
	detailsIdColumn = "details_id"
	createdAtColumn = "created_at"

	nameColumn            = "name"
	categoryColumn        = "category"
	transactionDateColumn = "transaction_date"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TransactionRepository {
	return &repo{db: db}
}

func (r *repo) CreateTransactionDetails(ctx context.Context, transactionDetails *model.TransactionDetailsInfo) (int64, error) {
	builder := sq.Insert(transactionDetailsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, categoryColumn, transactionDateColumn).
		Values(transactionDetails.Name, transactionDetails.Category, transactionDetails.TransactionDate).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.CreateTransactionDetails",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateTransaction(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsId int64) (int64, error) {
	builder := sq.Insert(transactionTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(ownerIdColumn, WalletIdColumn, boardIdColumn, sumColumn, detailsIdColumn).
		Values(transactionInfo.OwnerID, transactionInfo.WalletID, transactionInfo.BoardID, transactionInfo.Sum, transactionDetailsId).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.CreateTransaction",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
