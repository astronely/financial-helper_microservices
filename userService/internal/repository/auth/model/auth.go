package model

type User struct {
	ID   int64 `db:"id"`
	Info Info  `db:""`
}

type Info struct {
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
