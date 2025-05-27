package board

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *serv) SetBoard(ctx context.Context, boardID int64) error {
	err := s.CheckUserInBoardWithContext(ctx, boardID)
	if err != nil {
		logger.Error("error checking user in board | Service | SetBoard",
			"error", err.Error(),
		)
		return err
	}

	board, err := s.boardRepository.Get(ctx, boardID)
	if err != nil {
		logger.Error("error setting board id to cookie",
			"error", err.Error(),
		)
		return err
	}

	accessToken, err := utils.GenerateToken(
		board.ID,
		board.Info.OwnerID,
		[]byte(s.tokenConfig.AccessTokenKey()), s.tokenConfig.AccessTokenExpirationTime(),
	)

	if err != nil {
		logger.Error("error generating token in set board",
			"error", err.Error(),
		)
		return err
	}

	md := metadata.Pairs("set-cookie", "boardToken="+accessToken+"; HttpOnly; Path=/; Secure=false; SameSite=None")

	err = grpc.SendHeader(ctx, md)
	if err != nil {
		logger.Error("error send header in set board",
			"error", err.Error(),
		)
		return err
	}

	return nil
}
