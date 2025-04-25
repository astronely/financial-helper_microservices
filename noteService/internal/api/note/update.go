package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.service.Update(ctx, converter.ToNoteUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
