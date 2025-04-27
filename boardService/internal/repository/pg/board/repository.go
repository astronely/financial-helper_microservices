package board

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository"
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository/pg/board/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/boardService/internal/repository/pg/board/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

const (
	boardTableName     = "boards"
	boardUserTableName = "board_users"

	boardTableNameWithAlias     = "boards as b"
	boardUserTableNameWithAlias = "board_users as bu"

	boardPrefix     = "b."
	boardUserPrefix = "bu."

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	ownerIdColumn     = "owner"
	updatedAtColumn   = "updatedAt"
	createdAtColumn   = "createdAt"

	boardIdColumn = "board_id"
	userIdColumn  = "user_id"
	roleColumn    = "role"

	idColumnWithAlias             = boardPrefix + idColumn
	nameColumnWithAlias           = boardPrefix + nameColumn
	descriptionColumnWithAlias    = boardPrefix + descriptionColumn
	ownerIdColumnWithAlias        = boardPrefix + ownerIdColumn
	updatedAtColumnWithAlias      = boardPrefix + updatedAtColumn
	createdAtBoardColumnWithAlias = boardPrefix + createdAtColumn

	boardIdColumnWithAlias            = boardUserPrefix + boardIdColumn
	userIdColumnWithAlias             = boardUserPrefix + userIdColumn
	roleColumnWithAlias               = boardUserPrefix + roleColumn
	createdAtBoardUserColumnWithAlias = boardUserPrefix + createdAtColumn
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.BoardRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.BoardInfo) (int64, error) {
	builder := sq.Insert(boardTableName).
		Columns(nameColumn, descriptionColumn, ownerIdColumn).
		Values(info.Name, info.Description, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error create board | BuildToSql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "board_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error create board | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateUser(ctx context.Context, info *model.BoardUser) (int64, error) {
	builder := sq.Insert(boardUserTableName).
		Columns(boardIdColumn, userIdColumn, roleColumn).
		Values(info.BoardID, info.UserID, info.Role).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error create board user | BuildToSql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "board_repository.CreateUser",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error create board user | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Board, error) {
	builder := sq.Select(idColumn, nameColumn, descriptionColumn, ownerIdColumn, updatedAtColumn, createdAtColumn).
		From(boardTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error get board | BuildToSql",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "board_repository.Get",
		QueryRaw: query,
	}

	var board modelRepo.Board
	err = r.db.DB().ScanOneContext(ctx, &board, q, args...)
	if err != nil {
		logger.Error("error get board | ScanOneContext",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToBoardFromRepo(&board), nil
}

func (r *repo) GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error) {
	builder := sq.Select(boardIdColumn, userIdColumn, roleColumn, createdAtColumn).
		From(boardUserTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{boardIdColumn: boardId})

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error get board user | BuildToSql",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "board_repository.GetUser",
		QueryRaw: query,
	}

	var users []*modelRepo.BoardUser
	err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
	if err != nil {
		logger.Error("error get board user | ScanAllContext",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToBoardUserListFromRepo(users), nil
}

func (r *repo) ListByUserId(ctx context.Context, userId int64) ([]*model.Board, error) {
	usersBuilder := sq.Select(boardIdColumn).
		From(boardUserTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{userIdColumn: userId})

	query, args, err := usersBuilder.ToSql()
	if err != nil {
		logger.Error("error list by user id | BuildToSql userBuilder",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "board_repository.Users.ListByUserId",
		QueryRaw: query,
	}

	var ids []int64
	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		logger.Error("error list by user id | ScanAllContext userBuilder",
			"error", err.Error(),
		)
		return nil, err
	}

	boardsBuilder := sq.Select(idColumn, nameColumn, descriptionColumn, ownerIdColumn, updatedAtColumn, createdAtColumn).
		From(boardTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Expr("id IN (?)", ids))

	query, args, err = boardsBuilder.ToSql()
	if err != nil {
		logger.Error("error list by user id | BuildToSql boardsBuilder",
			"error", err.Error(),
		)
		return nil, err
	}

	q = db.Query{
		Name: "board_repository.Boards.ListByUserId",
	}

	var boards []*modelRepo.Board

	err = r.db.DB().ScanAllContext(ctx, &boards, q, args...)
	if err != nil {
		logger.Error("error list by user id | ScanAllContext boardsBuilder",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToBoardListFromRepo(boards), nil
}

func (r *repo) ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error) {
	builder := sq.Select(idColumn, nameColumn, descriptionColumn, ownerIdColumn, updatedAtColumn, createdAtColumn).
		From(boardTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{ownerIdColumn: ownerId})

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error list board | BuildToSql",
			"error", err.Error(),
		)
		return nil, err
	}

	q := db.Query{
		Name:     "board_repository.List",
		QueryRaw: query,
	}

	var boards []*modelRepo.Board
	err = r.db.DB().ScanAllContext(ctx, &boards, q, args...)
	if err != nil {
		logger.Error("error list board | ScanAllContext",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToBoardListFromRepo(boards), nil
}

func (r *repo) Update(ctx context.Context, info *model.BoardUpdate) (int64, error) {
	updateMap := make(map[string]interface{})
	if info.Name.Valid {
		updateMap[nameColumn] = info.Name.String
	}
	if info.Description.Valid {
		updateMap[descriptionColumn] = info.Description.String
	}

	builder := sq.Update(boardTableName).
		PlaceholderFormat(sq.Dollar).
		Table(boardTableName).
		SetMap(updateMap).
		Where(sq.Eq{idColumn: info.ID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error update board | BuildToSql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "board_repository.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error update board | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(boardTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error delete board | BuildToSql",
			"error", err.Error(),
		)
		return err
	}

	q := db.Query{
		Name:     "board_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("error delete board | ExecContext",
			"error", err.Error(),
		)
		return err
	}

	return nil
}

//func (r *repo) JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.JoinInfo, error) {
//	builder := sq.Insert(boardTableName).
//		PlaceholderFormat(sq.Dollar).
//		Columns(boardIdColumn, userIdColumn, roleColumn).
//		Values(info.BoarID, info.UserID, info.Role).
//		Suffix("RETURNING id")
//
//	query, args, err := builder.ToSql()
//	if err != nil {
//		logger.Error("error join board | BuildToSql",
//			"error", err.Error(),
//		)
//		return nil, err
//	}
//
//	q := db.Query{
//		Name:     "board_repository.Join",
//		QueryRaw: query,
//	}
//
//	var id int64
//	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
//	if err != nil {
//		logger.Error("error join board | QueryRawContext",
//			"error", err.Error(),
//		)
//		return nil, err
//	}
//
//	return info, nil
//}
