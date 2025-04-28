package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.service.Update(ctx, converter.ToBoardUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
