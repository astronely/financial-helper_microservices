package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.NoteUpdate) (int64, error) {
	id, err := s.noteRepository.Update(ctx, info)
	if err != nil {
		logger.Error("failed to update note",
			"error", err.Error(),
		)
		return 0, err
	}

	return id, nil
}
