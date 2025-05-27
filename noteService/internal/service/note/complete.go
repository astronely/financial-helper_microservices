package note

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	"github.com/astronely/financial-helper_microservices/noteService/internal/utils"
)

func (s *serv) Complete(ctx context.Context, info *model.NoteComplete) (int64, error) {
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

	if info.Status && !utils.CheckNoteOwner(ctx, userID, info.ID, s.noteRepository) &&
		!utils.CheckNotePerformer(ctx, userID, info.ID, s.noteRepository) &&
		!(userID == board.OwnerID) {
		return 0, errors.New("not authorized to complete note")
	}

	id, err := s.noteRepository.Complete(ctx, info, userID)
	if err != nil {
		logger.Error("failed to complete note",
			"id", info.ID,
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}
