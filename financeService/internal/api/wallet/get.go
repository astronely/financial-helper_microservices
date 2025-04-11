package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	wallet, err := i.service.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("Error getting wallet",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Get wallet",
		"wallet", wallet,
	)

	return &desc.GetResponse{
		Wallet: converter.ToWalletFromService(wallet),
	}, nil
}
