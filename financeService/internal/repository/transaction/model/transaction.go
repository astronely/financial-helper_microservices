package model

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID        int64           `db:"id"`
	Info      TransactionInfo `db:""`
	DetailsID int64           `db:"details_id"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`

	TransactionDetails TransactionDetails `db:""`
}

type TransactionInfo struct {
	OwnerID  int64           `db:"owner_id"`
	WalletID int64           `db:"wallet_id"`
	BoardID  int64           `db:"board_id"`
	Sum      decimal.Decimal `db:"sum"`
}

type TransactionDetails struct {
	ID              int64     `db:"detail_id"`
	Name            string    `db:"detail_name"`
	Category        int64     `db:"category"`
	TransactionDate time.Time `db:"transaction_date"`

	TransactionCategory TransactionCategory `db:""`
}

type TransactionCategory struct {
	ID          int64  `db:"id"`
	Name        string `db:"category_name"`
	Description string `db:"description"`
}

type TransactionInfoUpdate struct {
	WalletID int64           `db:"wallet_id"`
	SumDiff  decimal.Decimal `db:"sum_diff"`
}
