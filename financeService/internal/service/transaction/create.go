package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
)

func (s *serv) Create(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsInfo *model.TransactionDetailsInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		detailsId, errTx := s.transactionRepository.CreateTransactionDetails(ctx, transactionDetailsInfo)
		if errTx != nil {
			logger.Error("Failed to create transaction details",
				"error", errTx.Error(),
			)
			return errTx
		}

		id, errTx = s.transactionRepository.CreateTransaction(ctx, transactionInfo, detailsId)
		if errTx != nil {
			logger.Error("Failed to create transaction",
				"error", errTx.Error(),
			)
			return errTx
		}

		errTx = s.walletRepository.UpdateBalance(ctx, transactionInfo.WalletID, transactionInfo.Sum)
		if errTx != nil {
			logger.Error("Failed to update balance",
				"error", errTx.Error(),
			)
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
