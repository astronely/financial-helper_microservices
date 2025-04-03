package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/authService/internal/model"
	"github.com/astronely/financial-helper_microservices/authService/internal/repository"
	"github.com/astronely/financial-helper_microservices/authService/internal/repository/auth/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/authService/internal/repository/auth/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
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

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error) {
	builder := sq.Select(idColumn, emailColumn, nameColumn, passwordColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{emailColumn: email}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.get_user_by_email",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(&user), nil
}
