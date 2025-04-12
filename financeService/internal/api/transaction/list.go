package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	transactions, err := i.service.List(ctx, uint64(req.GetLimit()), uint64(req.GetOffset()))
	if err != nil {
		return nil, err
	}
	var convertedTransactions []*desc.GetResponse
	for _, transaction := range transactions {
		convertedTransactions = append(convertedTransactions, &desc.GetResponse{
			Transaction: converter.ToTransactionFromService(transaction),
			Details:     converter.ToTransactionDetailsFromService(transaction.TransactionDetails),
			Category:    converter.ToTransactionCategoryFromService(transaction.TransactionDetails.TransactionCategory),
		})
	}
	return &desc.ListResponse{
		Transactions: convertedTransactions,
	}, nil
}
