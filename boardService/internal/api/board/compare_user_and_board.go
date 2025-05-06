package board

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) CompareUserAndBoard(ctx context.Context, req *desc.CompareRequest) (*desc.CompareResponse, error) {
	result, err := i.service.CompareUserAndBoard(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.CompareResponse{
		Result: result,
	}, nil
}
