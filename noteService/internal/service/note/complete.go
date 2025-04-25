package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
)

func (s *serv) Complete(ctx context.Context, info *model.NoteComplete) (int64, error) {
	id, err := s.noteRepository.Complete(ctx, info)
	if err != nil {
		logger.Error("failed to complete note",
			"id", info.ID,
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}
