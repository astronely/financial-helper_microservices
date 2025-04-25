package note

import (
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/noteService/internal/service"
)

var _ def.NoteService = (*serv)(nil)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) def.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}
