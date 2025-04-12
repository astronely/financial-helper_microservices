package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) List(ctx context.Context, limit, offset uint64) ([]*model.Transaction, error) {
	transactions, err := s.transactionRepository.List(ctx, limit, offset)
	if err != nil {
		logger.Error("Error List transactions",
			"Error", err.Error())
		return nil, err
	}
	return transactions, nil
}
