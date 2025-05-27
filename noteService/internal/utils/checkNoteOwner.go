package utils

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
)

func CheckNoteOwner(ctx context.Context, userId int64, noteId int64, noteRepo repository.NoteRepository) bool {
	note, err := noteRepo.Get(ctx, noteId)
	if err != nil {
		logger.Error("get note in check note owner error",
			"error", err.Error(),
		)
		return false
	}

	ownerId := note.Info.OwnerID

	return ownerId == userId
}
