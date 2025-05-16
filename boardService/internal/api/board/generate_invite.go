package board

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) GenerateInvite(ctx context.Context, req *desc.GenerateInviteRequest) (*desc.GenerateInviteResponse, error) {
	url, err := i.service.GenerateInvite(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GenerateInviteResponse{
		Url: url,
	}, nil
}
