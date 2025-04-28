package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	board, err := i.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Board: converter.ToBoardFromService(board),
	}, nil
}
