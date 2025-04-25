package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

func (i *Implementation) Complete(ctx context.Context, req *desc.CompleteRequest) (*desc.CompleteResponse, error) {
	id, err := i.service.Complete(ctx, converter.ToNoteCompleteFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &desc.CompleteResponse{
		Id: id,
	}, nil
}
