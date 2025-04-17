package transaction

import (
	"context"
	"errors"
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

		// Create transaction in database
		id, errTx = s.transactionRepository.CreateTransaction(ctx, transactionInfo, detailsId)
		if errTx != nil {
			logger.Error("Failed to create transaction",
				"error", errTx.Error(),
			)
			return errTx
		}

		// Update wallet balance
		if transactionInfo.Type == "transfer" {
			errTx = s.walletRepository.UpdateBalance(ctx, transactionInfo.FromWalletID, transactionInfo.Amount.Neg(), "transfer")
			if errTx != nil {
				logger.Error("Failed to update balance",
					"info", "transfer, FromWalletID",
					"error", errTx.Error(),
				)
				return errTx
			}

			if !transactionInfo.ToWalletID.Valid {
				return errors.New("ToWalletID not valid")
			}
			errTx = s.walletRepository.UpdateBalance(ctx, transactionInfo.ToWalletID.Int64, transactionInfo.Amount, "transfer")
			if errTx != nil {
				logger.Error("Failed to update balance",
					"info", "transfer, ToWalletID",
					"error", errTx.Error(),
				)
				return errTx
			}
		} else {
			errTx = s.walletRepository.UpdateBalance(ctx, transactionInfo.FromWalletID, transactionInfo.Amount, "expense")
			if errTx != nil {
				logger.Error("Failed to update balance",
					"error", errTx.Error(),
				)
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
