package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) GenerateInviteURL(ctx context.Context, req *desc.GenerateInviteRequest) (*desc.GenerateInviteResponse, error) {
	url, err := i.service.GenerateInvite(ctx, converter.ToGenerateInviteFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.GenerateInviteResponse{
		Url: url,
	}, nil
}
