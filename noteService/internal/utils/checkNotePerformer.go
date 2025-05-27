package utils

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
)

func CheckNotePerformer(ctx context.Context, userId, noteId int64, noteRepo repository.NoteRepository) bool {
	note, err := noteRepo.Get(ctx, noteId)
	if err != nil {
		logger.Error("get note in check note performer error",
			"error", err.Error(),
		)
		return false
	}

	performerId := note.Info.PerformerID
	if !performerId.Valid {
		return false
	}

	return userId == performerId.Int64
}
