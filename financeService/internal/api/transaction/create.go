package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	//logger.Debug("Create impl",
	//	"get date", req.DetailsInfo.GetTransactionDate(),
	//	"req details date as time", req.DetailsInfo.GetTransactionDate().AsTime(),
	//)
	logger.Debug("Create transaction",
		"request", req)
	id, err := i.service.Create(ctx, converter.ToTransactionInfoFromDesc(req.GetInfo()),
		converter.ToTransactionDetailsInfoFromDesc(req.GetDetailsInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
