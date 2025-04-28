package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) GetUsers(ctx context.Context, req *desc.GetUsersRequest) (*desc.GetUsersResponse, error) {
	users, err := i.service.GetUsers(ctx, req.GetBoardId())
	if err != nil {
		return nil, err
	}

	return &desc.GetUsersResponse{
		Users: converter.ToBoardUsersFromService(users),
	}, nil
}
