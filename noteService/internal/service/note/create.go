package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.NoteCreate) (int64, error) {
	id, err := s.noteRepository.Create(ctx, info)
	if err != nil {
		logger.Error("failed to create note",
			"error", err.Error(),
		)
		return 0, err
	}
	return id, nil
}
