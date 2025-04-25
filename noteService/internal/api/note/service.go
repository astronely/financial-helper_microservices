package note

import (
	"github.com/astronely/financial-helper_microservices/noteService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	service service.NoteService
}

func NewImplementation(service service.NoteService) *Implementation {
	return &Implementation{
		service: service,
	}
}
