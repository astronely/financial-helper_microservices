package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/repository/user/converter"
	"github.com/astronely/financial-helper_microservices/internal/repository/user/modelRepo"
	_ "github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	tableName = "users"

	idColumn        = "id"
	emailColumn     = "email"
	nameColumn      = "name"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, nameColumn, passwordColumn).
		Values(info.Email, info.Name, password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, "", err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, "", err
	}
	return id, "token", nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	//log.Printf("TEST ID: %d", id)
	builder := sq.Select(idColumn, emailColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	//log.Printf("TEST QUERY: %s", query)
	if err != nil {
		return nil, err
	}

	var user modelRepo.User
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(&user), nil
}
