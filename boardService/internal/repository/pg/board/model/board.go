package model

import (
	"database/sql"
	"time"
)

type Board struct {
	ID        int64        `db:"id"`
	Info      BoardInfo    `db:""`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type BoardInfo struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	OwnerID     int64  `db:"owner_id"`
}

type BoardUser struct {
	BoardID   int64     `db:"board_id"`
	UserID    int64     `db:"user_id"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}
