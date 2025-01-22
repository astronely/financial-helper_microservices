package modelRepo

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Info      Info
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Info struct {
	Email string
	Name  string
}
