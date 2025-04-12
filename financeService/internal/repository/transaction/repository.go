package transaction

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

const (
	// Tables name
	transactionTableName        = "transactions"
	transactionDetailsTableName = "transaction_details"
	transactionCategoriesTable  = "transaction_categories"

	transactionTableNameWithAlias           = "transactions t"
	transactionDetailsTableNameWithAlias    = "transaction_details td"
	transactionCategoriesTableNameWithAlias = "transaction_categories tc"

	// Prefixes
	transactionPrefix           = "t."
	transactionDetailsPrefix    = "td."
	transactionCategoriesPrefix = "tc."

	// transactions
	// Основные колонки для таблицы "transactions"
	idColumn        = "id"
	ownerIdColumn   = "owner_id"
	walletIdColumn  = "wallet_id"
	boardIdColumn   = "board_id"
	sumColumn       = "sum"
	detailsIdColumn = "details_id"
	updatedAtColumn = "updated_at"
	createdAtColumn = "created_at"

	// Основные колонки для таблицы "transaction_details"
	detailsNameColumn     = "name"
	categoryColumn        = "category"
	transactionDateColumn = "transaction_date"

	// Основные колонки для таблицы "transaction_categories"
	categoryNameColumn = "name"
	descriptionColumn  = "description"

	// Основные колонки для таблицы "transactions"
	idColumnWithAlias        = transactionPrefix + idColumn
	ownerIdColumnWithAlias   = transactionPrefix + ownerIdColumn
	walletIdColumnWithAlias  = transactionPrefix + walletIdColumn
	boardIdColumnWithAlias   = transactionPrefix + boardIdColumn
	sumColumnWithAlias       = transactionPrefix + sumColumn
	detailsIdColumnWithAlias = transactionPrefix + detailsIdColumn
	updatedAtColumnWithAlias = transactionPrefix + updatedAtColumn
	createdAtColumnWithAlias = transactionPrefix + createdAtColumn

	// Основные колонки для таблицы "transaction_details"
	detailsNameColumnWithAlias     = transactionDetailsPrefix + detailsNameColumn + " AS detail_name"
	categoryColumnWithAlias        = transactionDetailsPrefix + categoryColumn
	transactionDateColumnWithAlias = transactionDetailsPrefix + transactionDateColumn

	// Основные колонки для таблицы "transaction_categories"
	categoryNameColumnWithAlias = transactionCategoriesPrefix + categoryNameColumn + " AS category_name"
	descriptionColumnWithAlias  = transactionCategoriesPrefix + descriptionColumn
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
		Columns(detailsNameColumn, categoryColumn, transactionDateColumn).
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
		Columns(ownerIdColumn, walletIdColumn, boardIdColumn, sumColumn, detailsIdColumn).
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

func (r *repo) Get(ctx context.Context, id int64) (*model.Transaction, error) {
	builder := sq.Select(
		idColumnWithAlias, ownerIdColumnWithAlias, walletIdColumnWithAlias, boardIdColumnWithAlias,
		sumColumnWithAlias, detailsIdColumnWithAlias, createdAtColumnWithAlias, updatedAtColumnWithAlias,
		detailsNameColumnWithAlias, categoryColumnWithAlias, transactionDateColumnWithAlias,
		categoryNameColumnWithAlias, descriptionColumnWithAlias,
	).
		PlaceholderFormat(sq.Dollar).
		From(transactionTableNameWithAlias).
		LeftJoin(transactionDetailsTableNameWithAlias + " ON t.details_id = td.id").
		LeftJoin(transactionCategoriesTableNameWithAlias + " ON td.category = tc.id").
		Where(sq.Eq{transactionPrefix + idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	//logger.Debug("SQL message from get",
	//	"SQL", query)

	q := db.Query{
		Name:     "finance_repository.Transaction.Get",
		QueryRaw: query,
	}

	var transaction modelRepo.Transaction
	err = r.db.DB().ScanOneContext(ctx, &transaction, q, args...)
	if err != nil {
		logger.Debug("SQL message from get",
			"Error", err.Error())
		return nil, err
	}

	return converter.ToTransactionFromRepo(&transaction), nil
}

func (r *repo) List(ctx context.Context, limit, offset uint64) ([]*model.Transaction, error) {
	builder := sq.Select(
		idColumnWithAlias, ownerIdColumnWithAlias, walletIdColumnWithAlias, boardIdColumnWithAlias,
		sumColumnWithAlias, detailsIdColumnWithAlias, createdAtColumnWithAlias, updatedAtColumnWithAlias,
		detailsNameColumnWithAlias, categoryColumnWithAlias, transactionDateColumnWithAlias,
		categoryNameColumnWithAlias, descriptionColumnWithAlias,
	).
		PlaceholderFormat(sq.Dollar).
		From(transactionTableNameWithAlias).
		LeftJoin(transactionDetailsTableNameWithAlias + " ON t.details_id = td.id").
		LeftJoin(transactionCategoriesTableNameWithAlias + " ON td.category = tc.id").
		Limit(limit).
		Offset(offset)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.List",
		QueryRaw: query,
	}

	var transactions []*modelRepo.Transaction
	err = r.db.DB().ScanAllContext(ctx, &transactions, q, args...)
	if err != nil {
		logger.Error("SQL Error message from list",
			"Error", err.Error())
		return nil, err
	}

	return converter.ToTransactionListFromRepo(transactions), nil
}
