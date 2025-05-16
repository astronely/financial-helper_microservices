package transaction

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
)

func (s *serv) Create(ctx context.Context, transactionInfo *model.TransactionInfo, transactionDetailsInfo *model.TransactionDetailsInfo) (int64, error) {
	var id int64
	//logger.Debug("Create transaction info",
	//	"transaction info", transactionInfo,
	//	"transaction details info", transactionDetailsInfo,
	//)
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
		if err != nil {
			logger.Error("error getting board id from context",
				"error", err.Error(),
			)
			return err
		}

		ownerID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
		if err != nil {
			logger.Error("error getting user id from context",
				"error", err.Error(),
			)
			return err
		}

		detailsId, errTx := s.transactionRepository.CreateTransactionDetails(ctx, transactionDetailsInfo)
		if errTx != nil {
			logger.Error("Failed to create transaction details",
				"error", errTx.Error(),
			)
			return errTx
		}

		transactionInfoFull := converter.AddOwnerAndBoardIdToTransactionInfo(transactionInfo, ownerID, boardID)
		// Create transaction in database
		id, errTx = s.transactionRepository.CreateTransaction(ctx, transactionInfoFull, detailsId)
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
			errTx = s.walletRepository.UpdateBalance(ctx, transactionInfo.FromWalletID, transactionInfo.Amount, transactionInfo.Type)
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
