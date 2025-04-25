package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int64        `db:"id"`
	Info      NoteInfo     `db:""`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type NoteInfo struct {
	BoardID        int64         `db:"board_id"`
	OwnerID        int64         `db:"owner_id"`
	PerformerID    sql.NullInt64 `db:"performer_id"`
	Content        string        `db:"content"`
	Status         bool          `db:"status"`
	CompletionDate sql.NullTime  `db:"completion_date"`
}
