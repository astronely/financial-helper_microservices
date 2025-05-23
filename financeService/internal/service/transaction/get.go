package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64, filters map[string]interface{}) (*model.Transaction, error) {
	transaction, err := s.transactionRepository.Get(ctx, id, filters)
	if err != nil {
		logger.Error("error getting transaction",
			"ID", id,
			"Error", err.Error(),
		)
		return nil, err
	}

	return transaction, nil
}
