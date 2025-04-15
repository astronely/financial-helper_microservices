package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	transaction, err := i.service.Get(ctx, req.GetId(), converter.Filters(req.GetFilterInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Transaction: converter.ToTransactionFromService(transaction),
		Details:     converter.ToTransactionDetailsFromService(transaction.TransactionDetails),
		Category:    converter.ToTransactionCategoryFromService(transaction.TransactionDetails.TransactionCategory),
	}, nil
}
