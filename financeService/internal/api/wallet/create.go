package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.service.Create(ctx, converter.ToCreateWalletInfoFromDesc(req.GetInfo()))
	if err != nil {
		logger.Error("Error creating wallet",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Creat wallet",
		"id", id,
	)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
