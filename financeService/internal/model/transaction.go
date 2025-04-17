package model

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID        int64
	Info      TransactionInfo
	DetailsId int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime

	TransactionDetails TransactionDetails
}

type TransactionInfo struct {
	OwnerID      int64
	FromWalletID int64
	ToWalletID   sql.NullInt64
	BoardID      int64
	Amount       decimal.Decimal
	Type         string
}

type TransactionDetails struct {
	ID   int64
	Info TransactionDetailsInfo

	TransactionCategory TransactionCategory
}

type TransactionDetailsInfo struct {
	Name            string
	Category        int64
	TransactionDate time.Time
}

type TransactionCategory struct {
	ID          int64
	Name        string
	Description string
}

type TransactionInfoUpdate struct {
	ID           int64
	FromWalletID int64
	ToWalletID   int64
	Amount       decimal.Decimal
	Type         string
}

type TransactionDetailsUpdate struct {
	ID       int64
	Name     string
	Category int64
}
