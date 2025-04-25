package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
)

func (s *serv) List(ctx context.Context, offset, limit uint64, filters map[string]interface{}) ([]*model.Note, error) {
	notes, err := s.noteRepository.List(ctx, offset, limit, filters)
	if err != nil {
		logger.Error("failed to list note",
			"error", err.Error(),
		)
		return nil, err
	}
	return notes, nil
}
