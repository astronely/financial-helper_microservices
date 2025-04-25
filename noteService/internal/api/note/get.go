package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := i.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		Note: converter.ToNoteFromService(note),
	}, nil
}
