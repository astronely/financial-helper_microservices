package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/repository/user/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/internal/repository/user/model"
	"github.com/astronely/financial-helper_microservices/pkg/client/db"
	_ "github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
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
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return 0, err
	}

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, nameColumn, passwordColumn).
		Values(info.Email, info.Name, hashedPassword).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, emailColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) List(ctx context.Context, limit uint64, offset uint64) ([]*model.User, error) {
	builder := sq.Select(idColumn, emailColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Limit(limit).
		Offset(offset)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.List",
		QueryRaw: query,
	}

	var users []*modelRepo.User
	err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUsersFromRepo(users), nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	return err
}

func (r *repo) Update(ctx context.Context, info *model.UpdateUserInfo) (int64, error) {
	updateMap := map[string]interface{}{}

	if info.Email != "" {
		updateMap[emailColumn] = info.Email
	}

	if info.Name != "" {
		updateMap[nameColumn] = info.Name
	}

	if info.Password != "" {
		updateMap[passwordColumn] = info.Password
	}

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		SetMap(updateMap).
		Where(sq.Eq{idColumn: info.ID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}
