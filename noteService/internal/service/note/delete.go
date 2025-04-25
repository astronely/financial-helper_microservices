package note

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.noteRepository.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete note",
			"id", id,
			"error", err.Error(),
		)
		return err
	}
	return nil
}
