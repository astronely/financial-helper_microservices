package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
	"google.golang.org/grpc/metadata"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	//wallets, err := i.service.List(ctx, uint64(req.GetLimit()), uint64(req.GetOffset()))
	md, ok := metadata.FromIncomingContext(ctx)
	logger.Debug("Metadata",
		"md", md,
		"ok", ok)
	wallets, err := i.service.List(ctx, req.GetBoardId())
	if err != nil {
		logger.Error("Error getting wallet",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Get wallet",
		"wallet", wallets,
	)

	return &desc.ListResponse{
		Wallets: converter.ToWalletListFromService(wallets),
	}, nil
}
