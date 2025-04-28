package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	boarId, err := i.service.Create(ctx, converter.ToBoardInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: boarId,
	}, nil
}
