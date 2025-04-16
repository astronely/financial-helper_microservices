package transaction

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"github.com/shopspring/decimal"
	"time"
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
	transactionDetailsIdWithAlias  = transactionDetailsPrefix + idColumn + " AS detail_id"
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

func (r *repo) Get(ctx context.Context, id int64, filters map[string]interface{}) (*model.Transaction, error) {
	builder := sq.Select(
		idColumnWithAlias, ownerIdColumnWithAlias, walletIdColumnWithAlias, boardIdColumnWithAlias,
		sumColumnWithAlias, detailsIdColumnWithAlias, updatedAtColumnWithAlias, createdAtColumnWithAlias,
		transactionDetailsIdWithAlias, detailsNameColumnWithAlias, categoryColumnWithAlias, transactionDateColumnWithAlias,
		categoryNameColumnWithAlias, descriptionColumnWithAlias,
	).
		PlaceholderFormat(sq.Dollar).
		From(transactionTableNameWithAlias).
		LeftJoin(transactionDetailsTableNameWithAlias + " ON t.details_id = td.id").
		LeftJoin(transactionCategoriesTableNameWithAlias + " ON td.category = tc.id").
		Where(sq.Eq{transactionPrefix + idColumn: id}).
		Limit(1)

	// Filtration block
	if val, ok := filters[categoryColumn]; ok {
		builder = builder.Where(sq.Eq{transactionDetailsPrefix + categoryColumn: val})
	}

	if val, ok := filters[transactionDateColumn]; ok {
		builder = builder.Where(sq.Eq{transactionDetailsPrefix + transactionDateColumn: val})
	}

	if val, ok := filters[ownerIdColumn]; ok {
		builder = builder.Where(sq.Eq{transactionPrefix + ownerIdColumn: val})
	}

	if val, ok := filters[walletIdColumn]; ok {
		builder = builder.Where(sq.Eq{transactionPrefix + walletIdColumn: val})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	//logger.Debug("SQL message from get before db.Query",
	//	"SQL", query)

	q := db.Query{
		Name:     "finance_repository.Transaction.Get",
		QueryRaw: query,
	}

	var transaction modelRepo.Transaction
	err = r.db.DB().ScanOneContext(ctx, &transaction, q, args...)
	if err != nil {
		logger.Error("Error getting transaction",
			"Error", err.Error())
		return nil, err
	}
	//logger.Debug("TransactionRepo",
	//	"transaction", transaction,
	//)
	return converter.ToTransactionFromRepo(&transaction), nil
}

func (r *repo) List(ctx context.Context, limit, offset uint64, filters map[string]interface{}) ([]*model.Transaction, error) {
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

	if val, ok := filters[categoryColumn]; ok {
		builder = builder.Where(sq.Eq{transactionDetailsPrefix + categoryColumn: val})
	}

	if val, ok := filters[transactionDateColumn]; ok {
		builder = builder.Where(sq.Eq{transactionDetailsPrefix + transactionDateColumn: val})
	}

	if val, ok := filters[ownerIdColumn]; ok {
		builder = builder.Where(sq.Eq{transactionPrefix + ownerIdColumn: val})
	}

	if val, ok := filters[walletIdColumn]; ok {
		builder = builder.Where(sq.Eq{transactionPrefix + walletIdColumn: val})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.List",
		QueryRaw: query,
	}
	//logger.Debug("SQL message from list before db.Query",
	//	"Filters", filters,
	//)
	//
	//logger.Debug("SQL message from list before db.Query",
	//	"SQL", query)

	var transactions []*modelRepo.Transaction
	err = r.db.DB().ScanAllContext(ctx, &transactions, q, args...)
	if err != nil {
		logger.Error("SQL Error message from list",
			"Error", err.Error())
		return nil, err
	}

	return converter.ToTransactionListFromRepo(transactions), nil
}

func (r *repo) UpdateInfo(ctx context.Context, updateInfo *model.TransactionInfoUpdate) (int64, decimal.Decimal, error) {
	updateTransactionInfoMap := make(map[string]interface{})

	if updateInfo.WalletID != 0 {
		updateTransactionInfoMap[walletIdColumn] = updateInfo.WalletID
	}
	if updateInfo.Sum.GreaterThanOrEqual(decimal.NewFromFloat(0)) {
		updateTransactionInfoMap[sumColumn] = updateInfo.Sum
	}
	updateTransactionInfoMap[updatedAtColumn] = time.Now()

	selectBuilder := sq.Select(sumColumn).
		From(transactionTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: updateInfo.ID}).
		Limit(1)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		logger.Error("SQL Error message from updateInfo",
			"Error", err.Error(),
		)
		return -1, decimal.Zero, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.UpdateInfo",
		QueryRaw: query,
	}

	var oldSum decimal.Decimal

	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&oldSum)
	if err != nil {
		logger.Error("SQL Error message from updateInfo from select",
			"Error", err.Error(),
		)
		return -1, decimal.Zero, err
	}

	updateBuilder := sq.Update(transactionTableName).
		PlaceholderFormat(sq.Dollar).
		SetMap(updateTransactionInfoMap).
		Where(sq.Eq{idColumn: updateInfo.ID}).
		Suffix(fmt.Sprintf("RETURNING %s", walletIdColumn))

	query, args, err = updateBuilder.ToSql()
	if err != nil {
		logger.Error("SQL Error message from updateInfo",
			"Error", err.Error(),
		)
		return -1, decimal.Zero, err
	}

	q = db.Query{
		Name:     "finance_repository.Transaction.UpdateInfo",
		QueryRaw: query,
	}

	var walletId int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&walletId)
	if err != nil {
		logger.Error("SQL Error message from updateInfo",
			"Error", err.Error(),
		)
		return -1, decimal.Zero, err
	}

	diff := updateInfo.Sum.Sub(oldSum)

	return walletId, diff, nil
}

func (r *repo) UpdateDetails(ctx context.Context, updateInfo *model.TransactionDetailsUpdate) (int64, error) {
	updateTransactionDetailsMap := make(map[string]interface{})

	if updateInfo.Name != "" {
		updateTransactionDetailsMap[detailsNameColumn] = updateInfo.Name
	}
	if updateInfo.Category != 0 {
		updateTransactionDetailsMap[categoryColumn] = updateInfo.Category
	}

	if len(updateTransactionDetailsMap) == 0 {
		return -1, errors.New("no update details found")
	}

	selectDetailsIdBuilder := sq.Select(detailsIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(transactionTableName).
		Where(sq.Eq{idColumn: updateInfo.ID}).
		Limit(1)

	query, args, err := selectDetailsIdBuilder.ToSql()
	if err != nil {
		logger.Error("SQL Error message from updateDetails",
			"Error", err.Error(),
		)
		return -1, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.UpdateDetails",
		QueryRaw: query,
	}

	var detailsId int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&detailsId)
	if err != nil {
		logger.Error("SQL Error message from updateDetails",
			"Error", err.Error(),
		)
		return -1, err
	}

	builder := sq.Update(transactionDetailsTableName).
		PlaceholderFormat(sq.Dollar).
		SetMap(updateTransactionDetailsMap).
		Where(sq.Eq{idColumn: detailsId}).
		Suffix("RETURNING id")

	query, args, err = builder.ToSql()
	if err != nil {
		logger.Error("SQL Error message from updateDetails",
			"Error", err.Error(),
		)
		return -1, err
	}

	q = db.Query{
		Name:     "finance_repository.Transaction.UpdateDetails",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("SQL Error message from updateDetails",
			"Error", err.Error(),
		)
		return -1, err
	}
	return id, nil
}

func (r *repo) Categories(ctx context.Context) ([]*model.TransactionCategory, error) {
	builder := sq.Select(idColumn, categoryNameColumn, descriptionColumn).
		PlaceholderFormat(sq.Dollar).
		From(transactionCategoriesTable)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("SQL Error message from categories",
			"Error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "finance_repository.Transaction.Categories",
		QueryRaw: query,
	}

	var categories []*model.TransactionCategory
	err = r.db.DB().ScanAllContext(ctx, &categories, q, args...)
	if err != nil {
		logger.Error("SQL Error message from categories",
			"Error", err.Error(),
		)
		return nil, err
	}

	return categories, nil
}
