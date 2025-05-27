package wallet

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
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

	wallet, err := s.walletRepository.Get(ctx, id)
	if err != nil {
		logger.Error("error getting wallet | Service | Delete",
			"error", err.Error(),
		)
		return err
	}

	if userId != board.OwnerID && userId != wallet.Info.OwnerID {
		return status.Error(codes.Unauthenticated, "not allowed")
	}

	err = s.walletRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("Failed to delete wallet",
			"error", err.Error())
	}
	return err
}
