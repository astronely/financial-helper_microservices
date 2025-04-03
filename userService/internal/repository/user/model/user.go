package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                int64          `db:"id"`
	Info              Info           `db:""`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         sql.NullTime   `db:"updated_at"`
	PasswordChangedAt sql.NullTime   `db:"password_changed_at"`
	Photo             sql.NullString `db:"photo"`
}

type Info struct {
	Email string `db:"email"`
	Name  string `db:"name"`
}
