package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	"github.com/astronely/financial-helper_microservices/noteService/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serv) Update(ctx context.Context, info *model.NoteUpdate) (int64, error) {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return 0, err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context",
			"error", err.Error(),
		)
		return 0, err
	}

	if !utils.CheckNoteOwner(ctx, userID, info.ID, s.noteRepository) &&
		!(userID == board.OwnerID) {
		return 0, status.Error(codes.Unauthenticated, "not allowed")
	}

	id, err := s.noteRepository.Update(ctx, info)
	if err != nil {
		logger.Error("failed to update note",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}
