package wallet

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
)

func (s *serv) Update(ctx context.Context, walletInfo *model.WalletUpdateInfo) (int64, error) {
	userId, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context | Service | Update",
			"error", err.Error(),
		)
		return 0, err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context | Service | Update",
			"error", err.Error(),
		)
		return 0, err
	}

	wallet, err := s.walletRepository.Get(ctx, walletInfo.ID)
	if err != nil {
		logger.Error("error getting wallet | Service | Delete",
			"error", err.Error(),
		)
		return 0, err
	}

	if userId != board.OwnerID && userId != wallet.Info.OwnerID {
		return 0, errors.New("not allowed")
	}

	id, err := s.walletRepository.Update(ctx, walletInfo)
	if err != nil {
		logger.Error("Failed to update wallet",
			"error", err.Error())
		return -1, err
	}
	return id, nil
}
