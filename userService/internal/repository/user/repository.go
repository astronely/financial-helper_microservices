package user

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/repository"
	"github.com/astronely/financial-helper_microservices/userService/internal/repository/user/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/userService/internal/repository/user/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	_ "github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	tableName = "users"

	idColumn                = "id"
	emailColumn             = "email"
	nameColumn              = "name"
	photoColumn             = "photo"
	passwordColumn          = "password"
	passwordChangedAtColumn = "password_changed_at"
	createdAtColumn         = "created_at"
	updatedAtColumn         = "updated_at"
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
		logger.Error("error create user | Builder",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error create user | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, emailColumn, nameColumn, createdAtColumn, updatedAtColumn, passwordChangedAtColumn, photoColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error get user | Builder",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		logger.Error("error get user | ScanOneContext",
			"error", err.Error(),
		)
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
		logger.Error("error list users | Builder",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.List",
		QueryRaw: query,
	}

	var users []*modelRepo.User
	err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
	if err != nil {
		logger.Error("error list users | ScanAllContext",
			"error", err.Error(),
		)
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

	result, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		logger.Error("error delete user | ExecContext",
			"error", err.Error(),
		)
		return err
	}
	if result.RowsAffected() == 0 {
		logger.Error("error delete user | RowsAffected = 0",
			"error", errors.New("no user to delete"),
		)
		return errors.New("no user to delete")
	}
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

	updateMap["updated_at"] = time.Now()

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		SetMap(updateMap).
		Where(sq.Eq{idColumn: info.ID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		logger.Error("error update user | Builder",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error update user | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, err
}
