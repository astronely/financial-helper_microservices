package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.service.Create(ctx, converter.ToNoteCreateFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
