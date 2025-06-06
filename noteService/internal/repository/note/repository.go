package note

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository/note/converter"
	modelRepo "github.com/astronely/financial-helper_microservices/noteService/internal/repository/note/model"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"time"
)

const (
	tableName = "notes"

	idColumn             = "id"
	boardIdColumn        = "board_id"
	ownerIdColumn        = "owner_id"
	performerIdColumn    = "performer_id"
	contentColumn        = "content"
	statusColumn         = "status"
	completionDateColumn = "completion_date"
	updatedAtColumn      = "updated_at"
	createdAtColumn      = "created_at"

	// Filters columns
	completionDateStartColumn = "completed_start"
	completionDateEndColumn   = "completed_end"
	createdAtStartColumn      = "created_at_start"
	createdAtEndColumn        = "created_at_end"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.NoteRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.NoteCreate) (int64, error) {
	logger.Debug("Info in repo create",
		"info", info)
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(boardIdColumn, ownerIdColumn, contentColumn).
		Values(info.BoardID, info.OwnerID, info.Content).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error create note | build to sql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "note_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error create note | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select(idColumn, ownerIdColumn, performerIdColumn, contentColumn, statusColumn, completionDateColumn, updatedAtColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error get note | BuildToSql",
			"error", err.Error(),
		)
		return nil, err
	}
	q := db.Query{
		Name:     "note_repository.Get",
		QueryRaw: query,
	}

	logger.Debug("get note from database",
		"query", query)
	var note modelRepo.Note
	err = r.db.DB().ScanOneContext(ctx, &note, q, args...)
	if err != nil {
		logger.Error("error get note | ScanOneContext",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToNoteFromRepo(&note), nil
}
func (r *repo) List(ctx context.Context, boardID int64, limit, offset uint64, filters map[string]interface{}) ([]*model.Note, error) {
	builder := sq.Select(idColumn, ownerIdColumn, performerIdColumn, contentColumn, statusColumn, completionDateColumn, updatedAtColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{boardIdColumn: boardID}).
		OrderBy(createdAtColumn + " DESC")

	if limit > 0 {
		builder = builder.Limit(limit)
	}

	if offset > 0 {
		builder = builder.Offset(offset)
	}

	logger.Debug("filters",
		"filters", filters[ownerIdColumn])

	if val, ok := filters[ownerIdColumn]; ok {
		logger.Debug("filters, ownerId",
			"value", val)
		builder = builder.Where(sq.Eq{ownerIdColumn: val})
	}
	if val, ok := filters[performerIdColumn]; ok {
		logger.Debug("filters, performerId",
			"value", val)
		builder = builder.Where(sq.Eq{performerIdColumn: val})
	}
	if val, ok := filters[statusColumn]; ok {
		logger.Debug("filters, status",
			"value", val)
		builder = builder.Where(sq.Eq{statusColumn: val})
	}

	if val, ok := filters[createdAtStartColumn]; ok {
		logger.Debug("filters, created_at_start",
			"value", val,
		)
		builder = builder.Where(sq.GtOrEq{createdAtColumn: val})
	}
	if val, ok := filters[createdAtEndColumn]; ok {
		builder = builder.Where(sq.LtOrEq{createdAtColumn: val})
	}
	if val, ok := filters[completionDateStartColumn]; ok {
		builder = builder.Where(sq.GtOrEq{completionDateColumn: val})
	}
	if val, ok := filters[completionDateEndColumn]; ok {
		builder = builder.Where(sq.LtOrEq{completionDateColumn: val})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error list note | BuildToSql",
			"error", err.Error(),
		)
		return nil, err
	}
	logger.Debug("list query",
		"query", query,
	)
	q := db.Query{
		Name:     "note_repository.List",
		QueryRaw: query,
	}

	var notes []*modelRepo.Note
	err = r.db.DB().ScanAllContext(ctx, &notes, q, args...)
	if err != nil {
		logger.Error("error list note | ScanAllContext",
			"error", err.Error(),
		)
		return nil, err
	}

	return converter.ToNoteListFromRepo(notes), nil
}
func (r *repo) Update(ctx context.Context, info *model.NoteUpdate) (int64, error) {
	updateDate := time.Now()
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(contentColumn, info.Content).
		Set(updatedAtColumn, updateDate).
		Where(sq.Eq{idColumn: info.ID}).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error update note | BuildToSql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "note_repository.Update",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error update note | QueryRawContext",
			"error", err.Error(),
		)
		return 0, err
	}
	return id, nil
}
func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error delete note | BuildToSql",
			"error", err.Error(),
		)
		return err
	}

	q := db.Query{
		Name:     "note_repository.Delete",
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error("error delete note | ExecContext",
			"error", err.Error(),
		)
		return err
	}

	if result.RowsAffected() == 0 {
		logger.Error("error delete note | RowsAffected = 0",
			"error", errors.New("no note to delete"),
		)
		return errors.New("no note to delete")
	}

	return nil
}
func (r *repo) Complete(ctx context.Context, info *model.NoteComplete, performerID int64) (int64, error) {
	updateDate := time.Now()
	var completionDate time.Time
	var performerIdToSet sql.NullInt64
	if info.Status {
		completionDate = time.Now()
		performerIdToSet = sql.NullInt64{
			Int64: performerID,
			Valid: true,
		}
	}

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(statusColumn, info.Status).
		Set(completionDateColumn, completionDate).
		Set(updatedAtColumn, updateDate).
		Set(performerIdColumn, performerIdToSet).
		Where(sq.Eq{idColumn: info.ID}).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		logger.Error("error complete note | BuildToSql",
			"error", err.Error(),
		)
		return 0, err
	}

	q := db.Query{
		Name:     "note_repository.Complete",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRawContext(ctx, q, args...).Scan(&id)
	if err != nil {
		logger.Error("error complete note | QueryRawContext",
			"error", err.Error(),
		)
	}

	return id, nil
}
