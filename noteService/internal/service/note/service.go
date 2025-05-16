package note

import (
	"github.com/astronely/financial-helper_microservices/noteService/internal/config"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/noteService/internal/service"
)

var _ def.NoteService = (*serv)(nil)

type serv struct {
	noteRepository repository.NoteRepository

	tokenConfig config.TokenConfig
}

func NewService(noteRepository repository.NoteRepository, tokenConfig config.TokenConfig) def.NoteService {
	return &serv{
		noteRepository: noteRepository,
		tokenConfig:    tokenConfig,
	}
}
