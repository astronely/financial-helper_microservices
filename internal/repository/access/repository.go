package access

import (
	"context"
	_ "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/internal/client/db"
)

const (
	tableName = "users"

	idColumn       = "id"
	emailColumn    = "email"
	nameColumn     = "name"
	passwordColumn = "password"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CheckUser(ctx context.Context, id int64) (bool, error) {
	//builder := sq.Select(idColumn, emailColumn, nameColumn, passwordColumn).
	//	From(tableName).
	//	Where(sq.Eq{idColumn: id}).
	//	Limit(1)
	//
	//query, args, err := builder.ToSql()
	//
	//if err != nil {
	//	return false, err
	//}
	//
	//q := db.Query{
	//	Name: "access_repository.check_user",
	//	QueryRaw: query,
	//}
	//
	//var user
	return false, nil
}
