package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	transactions, err := i.service.List(ctx, req.GetBoardId(), uint64(req.GetLimit()), uint64(req.GetOffset()),
		converter.Filters(req.GetFilterInfo()))

	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{
		Transactions: converter.ToTransactionListFromService(transactions),
	}, nil
}
