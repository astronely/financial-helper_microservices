package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
)

func (s *serv) Create(ctx context.Context, walletInfo *model.CreateWalletInfo) (int64, error) {
	boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board id",
			"error", err.Error(),
		)
		return 0, err
	}

	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id",
			"error", err.Error(),
		)
		return 0, err
	}

	walletInfoFull := converter.AddOwnerAndBoardIdToWalletInfo(walletInfo, userID, boardID)

	id, err := s.walletRepository.Create(ctx, walletInfoFull)
	if err != nil {
		logger.Error("Failed to create new wallet",
			"error", err.Error())
		return 0, err
	}
	return id, nil
}
