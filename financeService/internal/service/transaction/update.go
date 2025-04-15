package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/shopspring/decimal"
)

func (s *serv) Update(ctx context.Context,
	updateInfo *model.TransactionInfoUpdate,
	updateDetails *model.TransactionDetailsUpdate) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var walletId int64
		var diff decimal.Decimal
		var errTx error

		if updateInfo.WalletID != 0 || !updateInfo.Sum.Equal(decimal.NewFromInt(-1)) {
			walletId, diff, errTx = s.transactionRepository.UpdateInfo(ctx, updateInfo)
			if errTx != nil {
				logger.Error("error updating transaction info",
					"Error", errTx,
				)
				return errTx
			}

			if !updateInfo.Sum.Equal(decimal.NewFromInt(-1)) {
				errTx = s.walletRepository.UpdateBalance(ctx, walletId, diff)
				if errTx != nil {
					logger.Error("error updating wallet balance after changing transaction info",
						"Error", errTx,
					)
					return errTx
				}
			}
		}

		id, errTx = s.transactionRepository.UpdateDetails(ctx, updateDetails)
		if errTx != nil {
			logger.Error("error updating details",
				"Error", errTx,
			)
			return errTx
		}
		return nil
	})

	if err != nil {
		logger.Error("error updating transaction info",
			"Error", err.Error(),
		)
		return 0, err
	}
	return id, nil
}
