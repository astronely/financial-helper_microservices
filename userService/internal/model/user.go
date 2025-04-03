package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Email string
	Name  string
}

type UpdateUserInfo struct {
	ID       int64
	Email    string
	Name     string
	Password string
}
