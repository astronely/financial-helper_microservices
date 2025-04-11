package model

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	ID        int64        `db:"id"`
	Info      WalletInfo   `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type WalletInfo struct {
	OwnerID int64           `db:"owner_id"`
	BoardID int64           `db:"board_id"`
	Name    string          `db:"name"`
	Balance decimal.Decimal `db:"balance"`
}
