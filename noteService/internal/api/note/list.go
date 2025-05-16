package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	notes, err := i.service.List(ctx, req.GetBoardId(), uint64(req.GetLimit()), uint64(req.GetOffset()), converter.ToFilters(req.GetFilterInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{
		Notes: converter.ToNoteListFromService(notes),
	}, nil
}
