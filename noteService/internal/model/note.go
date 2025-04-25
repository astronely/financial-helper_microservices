package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int64
	Info      NoteInfo
	UpdatedAt sql.NullTime
	CreatedAt time.Time
}

type NoteInfo struct {
	BoardID        int64
	OwnerID        int64
	PerformerID    sql.NullInt64
	Content        string
	Status         bool
	CompletionDate sql.NullTime
}

type NoteCreate struct {
	BoardID int64
	OwnerID int64
	Content string
}

type NoteUpdate struct {
	ID      int64
	Content string
}

type NoteComplete struct {
	ID     int64
	Status bool
}
