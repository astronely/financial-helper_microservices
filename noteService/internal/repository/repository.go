package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	List(ctx context.Context, boardID int64, offset, limit uint64, filters map[string]interface{}) ([]*model.Note, error)
	Update(ctx context.Context, info *model.NoteUpdate) (int64, error)
	Delete(ctx context.Context, id int64) error
	Complete(ctx context.Context, info *model.NoteComplete, performerID int64) (int64, error)
}
