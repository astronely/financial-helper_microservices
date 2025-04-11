package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.service.Delete(ctx, req.GetId())
	if err != nil {
		logger.Error("Error deleting wallet",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Delete wallet",
		"id", req.GetId(),
	)

	return &emptypb.Empty{}, nil
}
