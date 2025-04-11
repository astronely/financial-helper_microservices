package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.service.Update(ctx, converter.ToUpdateWalletInfoFromDesc(req))
	if err != nil {
		logger.Error("Error getting wallet",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Update wallet",
		"id", id,
	)

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
