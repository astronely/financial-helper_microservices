package model

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	ID        int64
	Info      WalletInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime // В БД надо добавить это поле, его нет
}

type CreateWalletInfo struct {
	OwnerID int64
	BoardID int64
	Name    string
}

type WalletInfo struct {
	OwnerID int64
	BoardID int64
	Name    string
	Balance decimal.Decimal
}

type WalletUpdateInfo struct {
	ID      int64
	Name    string
	Balance decimal.Decimal
}
