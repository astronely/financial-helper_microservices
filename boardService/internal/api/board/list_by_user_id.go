package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) ListByUserId(ctx context.Context, req *desc.ListByUserIdRequest) (*desc.ListResponse, error) {
	boards, err := i.service.ListByUserId(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{
		Boards: converter.ToBoardListFromService(boards),
	}, nil
}
