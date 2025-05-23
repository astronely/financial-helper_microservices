package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) List(ctx context.Context, boardID int64, limit, offset uint64, filters map[string]interface{}) ([]*model.Transaction, error) {
	transactions, err := s.transactionRepository.List(ctx, boardID, limit, offset, filters)
	if err != nil {
		logger.Error("Error List transactions",
			"Error", err.Error())
		return nil, err
	}
	return transactions, nil
}
