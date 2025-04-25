package converter

import (
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/noteService/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {
	return &model.Note{
		ID:        note.ID,
		Info:      ToNoteInfoFromRepo(note.Info),
		UpdatedAt: note.UpdatedAt,
		CreatedAt: note.CreatedAt,
	}
}

func ToNoteInfoFromRepo(noteInfo modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		BoardID:        noteInfo.BoardID,
		OwnerID:        noteInfo.OwnerID,
		PerformerID:    noteInfo.PerformerID,
		Content:        noteInfo.Content,
		Status:         noteInfo.Status,
		CompletionDate: noteInfo.CompletionDate,
	}
}

func ToNoteListFromRepo(notes []*modelRepo.Note) []*model.Note {
	var noteList []*model.Note
	for _, note := range notes {
		noteList = append(noteList, ToNoteFromRepo(note))
	}
	return noteList
}
