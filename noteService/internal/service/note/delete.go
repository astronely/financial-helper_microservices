package note

import (
	"context"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/utils"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return err
	}

	board, err := utils.GetBoardFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board from context",
			"error", err.Error(),
		)
		return err
	}

	if !utils.CheckNoteOwner(ctx, userID, id, s.noteRepository) &&
		!(userID == board.OwnerID) {
		return errors.New("not authorized to complete note")
	}

	err = s.noteRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete note",
			"id", id,
			"error", err.Error(),
		)
		return err
	}
	return nil
}
