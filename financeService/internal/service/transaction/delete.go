package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serv) Delete(ctx context.Context, transactionId int64) error {
	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	transaction, err := s.transactionRepository.Get(ctx, transactionId, map[string]interface{}{})
	if err != nil {
		logger.Error("error getting wallet | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	if userId != board.OwnerID && userId != transaction.Info.OwnerID {
		return status.Error(codes.Unauthenticated, "not allowed")
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		var transactionInfo *model.Transaction

		transactionInfo, errTx = s.transactionRepository.Get(ctx, transactionId, nil)
		if errTx != nil {
			logger.Error("Failed to get transaction in DeleteTx",
				"error", errTx.Error(),
			)
			return errTx
		}

		errTx = s.transactionRepository.Delete(ctx, transactionId)
		if errTx != nil {
			logger.Error("Failed to delete transaction in DeleteTx",
				"error", errTx.Error(),
			)
			return errTx
		}

		var amountToDecrease = transactionInfo.Info.Amount.Neg()
		if transactionInfo.Info.Type == "transfer" {
			amountToDecrease = amountToDecrease.Neg()
		}

		errTx = s.walletRepository.UpdateBalance(ctx,
			transactionInfo.Info.FromWalletID,
			amountToDecrease,
			transactionInfo.Info.Type,
		)
		if errTx != nil {
			logger.Error("Failed to update transaction in DeleteTx",
				"error", errTx.Error(),
			)
			return errTx
		}

		if transactionInfo.Info.ToWalletID.Valid {
			// If ToWalletID valid -> type = transfer
			// If transfer -> updateBalance use +amount
			errTx = s.walletRepository.UpdateBalance(ctx,
				transactionInfo.Info.ToWalletID.Int64,
				transactionInfo.Info.Amount.Neg(),
				transactionInfo.Info.Type,
			)
			if errTx != nil {
				logger.Error("Failed to update transaction in DeleteTx",
					"error", errTx.Error(),
				)
				return errTx
			}
		}
		return nil
	})

	if err != nil {
		logger.Error("Failed to delete transaction",
			"error", err.Error(),
		)
		return err
	}
	return nil
}
