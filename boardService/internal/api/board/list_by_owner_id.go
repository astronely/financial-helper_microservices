package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) ListByOwnerId(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	boards, err := i.service.ListByOwnerId(ctx, req.GetOwnerId())
	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{
		Boards: converter.ToBoardListFromService(boards),
	}, nil
}
