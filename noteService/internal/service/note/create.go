package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/converter"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	"github.com/astronely/financial-helper_microservices/noteService/internal/utils"
)

func (s *serv) Create(ctx context.Context, info *model.NoteCreate) (int64, error) {
	boardID, err := utils.GetBoardIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting board id from context",
			"error", err.Error(),
		)
		return 0, err
	}

	userID, err := utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
	if err != nil {
		logger.Error("error getting user id from context",
			"error", err.Error(),
		)
		return 0, err
	}

	infoFull := converter.AddOwnerAndBoardIdToNoteCreate(info, userID, boardID)

	id, err := s.noteRepository.Create(ctx, infoFull)
	if err != nil {
		logger.Error("failed to create note",
			"error", err.Error(),
		)
		return 0, err
	}
	return id, nil
}
