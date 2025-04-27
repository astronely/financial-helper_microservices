package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) Join(ctx context.Context, req *desc.JoinRequest) (*desc.JoinResponse, error) {
	info, err := i.service.JoinBoard(ctx, converter.ToJoinInfoFromDesc(req))

	if err != nil {
		return nil, err
	}

	logger.Debug("Token",
		"token", req.GetToken())

	return &desc.JoinResponse{
		Id: info.BoarID,
	}, nil
}
