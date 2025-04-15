package wallet

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository/wallet/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/wallet/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"github.com/shopspring/decimal"
	"time"
)

const (
	tableName = "wallets"

	idColumn        = "id"
	ownerIdColumn   = "owner_id"
	boardIdColumn   = "board_id"
	nameColumn      = "name"
	balanceColumn   = "balance"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.WalletRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, walletInfo *model.WalletInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(ownerIdColumn, boardIdColumn, nameColumn, balanceColumn).
		Values(walletInfo.OwnerID, walletInfo.BoardID, walletInfo.Name, walletInfo.Balance).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "finance_repository.Wallet.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Wallet, error) {
	builder := sq.Select(idColumn, ownerIdColumn, boardIdColumn, nameColumn, balanceColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "finance_repository.Wallet.Get",
		QueryRaw: query,
	}

	var wallet modelRepo.Wallet
	err = r.db.DB().ScanOneContext(ctx, &wallet, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToWalletFromRepo(&wallet), nil
}

func (r *repo) List(ctx context.Context, limit, offset uint64) ([]*model.Wallet, error) {
	builder := sq.Select(idColumn, ownerIdColumn, boardIdColumn, nameColumn, balanceColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Limit(limit).
		Offset(offset)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "finance_repository.Wallet.List",
		QueryRaw: query,
	}

	var wallets []*modelRepo.Wallet
	err = r.db.DB().ScanAllContext(ctx, &wallets, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToWalletsFromRepo(wallets), nil
}

func (r *repo) Update(ctx context.Context, walletInfo *model.WalletUpdateInfo) (int64, error) {
	updateMap := map[string]interface{}{}

	if walletInfo.Name != "" {
		updateMap[nameColumn] = walletInfo.Name
	}

	logger.Debug("Debug balance",
		"walletInfo balance", walletInfo.Balance,
		"decimal to check", decimal.NewFromInt(-1),
		"if equal", walletInfo.Balance.Equal(decimal.NewFromInt(-1)))

	if !walletInfo.Balance.Equal(decimal.NewFromInt(-1)) {
		updateMap[balanceColumn] = walletInfo.Balance
	}

	updateMap[updatedAtColumn] = time.Now()

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		SetMap(updateMap).
		Where(sq.Eq{idColumn: walletInfo.ID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	logger.Debug("Wallet args",
		"args", args,
	)

	q := db.Query{
		Name:     "finance_repository.Wallet.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) UpdateBalance(ctx context.Context, id int64, value decimal.Decimal) error {
	//logger.Debug("Debug balance",
	//	"id", id,
	//	"Value to add", value,
	//)
	builder := sq.Select(balanceColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "finance_repository.Wallet.UpdateBalance.GetBalance",
		QueryRaw: query,
	}
	var balance decimal.Decimal

	err = r.db.DB().ScanOneContext(ctx, &balance, q, args...)
	if err != nil {
		return err
	}

	balance = balance.Add(value)

	if balance.IsNegative() {
		return errors.New("negative balance")
	}

	builder2 := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(balanceColumn, balance).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err = builder2.ToSql()
	if err != nil {
		return err
	}

	q = db.Query{
		Name:     "finance_repository.Wallet.UpdateBalance",
		QueryRaw: query,
	}

	//logger.Debug("Debug balance",
	//	"query", q,
	//	"balance", balance,
	//)

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "finance_repository.Wallet.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if res.RowsAffected() == 0 {
		return errors.New("wallet with this id doesnt exist")
	}

	return err
}
