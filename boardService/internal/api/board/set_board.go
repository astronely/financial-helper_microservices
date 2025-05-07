package board

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) SetBoard(ctx context.Context, req *desc.SetBoardRequest) (*desc.SetBoardResponse, error) {
	err := i.service.SetBoard(ctx, req.GetBoardId())
	if err != nil {
		return nil, err
	}

	return &desc.SetBoardResponse{
		Result: true,
	}, nil
}
