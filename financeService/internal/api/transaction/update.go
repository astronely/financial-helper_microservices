package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.service.Update(ctx,
		converter.ToTransactionInfoUpdateFromDesc(req),
		converter.ToTransactionDetailsUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
