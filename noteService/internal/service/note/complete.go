package note

import (
	"context"
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
