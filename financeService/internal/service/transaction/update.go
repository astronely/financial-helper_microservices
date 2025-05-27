package transaction

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
	"github.com/shopspring/decimal"
)

func (s *serv) Update(ctx context.Context,
	updateInfo *model.TransactionInfoUpdate,
	updateDetails *model.TransactionDetailsUpdate) (int64, error) {

	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Delete",
			"error", err.Error(),
		)
		return 0, err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context | Service | Delete",
			"error", err.Error(),
		)
		return 0, err
	}

	transaction, err := s.transactionRepository.Get(ctx, updateInfo.ID, map[string]interface{}{})
	if err != nil {
		logger.Error("error getting wallet | Service | Delete",
			"error", err.Error(),
		)
		return 0, err
	}

	if userId != board.OwnerID && userId != transaction.Info.OwnerID {
		return 0, errors.New("not allowed")
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		if updateInfo.ToWalletID != 0 && updateInfo.Type != "transfer" {
			return errors.New("to wallet id must be with type transfer")
		}
		if updateInfo.Type == "transfer" && updateInfo.ToWalletID == 0 {
			return errors.New("type transfer must be with wallet id")
		}

		var oldTxInfo *model.TransactionInfo
		var errTx error

		// Update info if new info exists
		if updateInfo.FromWalletID != 0 || !updateInfo.Amount.Equal(decimal.NewFromInt(-1)) ||
			updateInfo.ToWalletID != 0 || updateInfo.Type != "" {

			oldTxInfo, errTx = s.transactionRepository.UpdateInfo(ctx, updateInfo)
			if errTx != nil {
				logger.Error("error updating transaction info",
					"Error", errTx.Error(),
				)
				return errTx
			}

			// Merge New and Old info
			var amount = oldTxInfo.Amount
			var txType = oldTxInfo.Type
			var fromWalletId = oldTxInfo.FromWalletID
			var toWalletId = oldTxInfo.ToWalletID.Int64

			if !updateInfo.Amount.Equal(decimal.NewFromInt(-1)) {
				amount = updateInfo.Amount
			}
			if updateInfo.Type != "" {
				txType = updateInfo.Type
			}
			if updateInfo.FromWalletID != 0 {
				fromWalletId = updateInfo.FromWalletID
			}
			if updateInfo.ToWalletID != 0 {
				toWalletId = updateInfo.ToWalletID
			}

			// If new info exists, do rollback and add operations
			if !updateInfo.Amount.Equal(decimal.NewFromInt(-1)) || updateInfo.FromWalletID != 0 ||
				updateInfo.ToWalletID != 0 || updateInfo.Type != "" {

				// Rollback FromWalletId (old info)
				amountToAdd := oldTxInfo.Amount.Neg()
				if oldTxInfo.Type == "transfer" {
					amountToAdd = amountToAdd.Neg()
				}

				errTx = s.walletRepository.UpdateBalance(ctx, oldTxInfo.FromWalletID, amountToAdd, oldTxInfo.Type)
				if errTx != nil {
					logger.Error("error updating transaction balance, rollback from wallet id balance",
						"Error", errTx.Error(),
					)
					return errTx
				}

				// Add amount to FromWalletID (new info)
				amountToAdd = amount
				if updateInfo.Type == "transfer" {
					amountToAdd = amountToAdd.Neg()
				}

				errTx = s.walletRepository.UpdateBalance(ctx, fromWalletId, amountToAdd, txType)
				if errTx != nil {
					logger.Error("error updating wallet balance, FromWalletId",
						"Error", errTx.Error(),
					)
					return errTx
				}

				// Rollback ToWalletId
				if oldTxInfo.ToWalletID.Valid {
					// If ToWalletID exists, then type = transfer
					errTx = s.walletRepository.UpdateBalance(ctx, oldTxInfo.ToWalletID.Int64, oldTxInfo.Amount.Neg(), oldTxInfo.Type)
					if errTx != nil {
						logger.Error("error updating wallet balance, rollback to wallet id balance",
							"Error", errTx.Error(),
						)
						return errTx
					}
				}

				// Add amount to ToWalletID (new info)
				if updateInfo.ToWalletID != 0 {
					errTx = s.walletRepository.UpdateBalance(ctx, toWalletId, amount, txType)
					if errTx != nil {
						logger.Error("error updating wallet balance, add ToWalletId",
							"Error", errTx.Error(),
						)
					}
				}
			}
		}

		// Update details if new details exists
		if updateDetails.Name != "" || updateDetails.Category != 0 {
			_, errTx = s.transactionRepository.UpdateDetails(ctx, updateDetails)
			if errTx != nil {
				logger.Error("error updating details",
					"Error", errTx,
				)
				return errTx
			}
		}
		return nil
	})

	if err != nil {
		logger.Error("error updating transaction info",
			"Error", err.Error(),
		)
		return 0, err
	}
	return updateInfo.ID, nil
}
