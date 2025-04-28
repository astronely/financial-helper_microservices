package model

import (
	"database/sql"
	"time"
)

type Board struct {
	ID        int64
	Info      BoardInfo
	UpdatedAt sql.NullTime
	CreatedAt time.Time
}

type BoardInfo struct {
	Name        string
	Description string
	OwnerID     int64
}

type BoardUser struct {
	BoardID   int64
	UserID    int64
	Role      string
	CreatedAt time.Time
}

type BoardUserCreate struct {
	BoardID int64
	UserID  int64
	Role    string
}

type BoardUpdate struct {
	ID          int64
	Name        sql.NullString
	Description sql.NullString
}

type JoinInfo struct {
	Token string
	ID    int64
}

type GenerateInviteInfo struct {
	BoardID int64
	//UserID  int64
	Role string
}
