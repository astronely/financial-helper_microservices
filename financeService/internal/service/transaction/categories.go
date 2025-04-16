package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Categories(ctx context.Context) ([]*model.TransactionCategory, error) {
	categories, err := s.transactionRepository.Categories(ctx)
	if err != nil {
		logger.Error("error getting categories",
			"error", err.Error(),
		)
		return nil, err
	}

	return categories, nil
}
