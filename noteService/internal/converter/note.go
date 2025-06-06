package converter

import (
	"database/sql"
	"github.com/astronely/financial-helper_microservices/noteService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/noteService/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	idColumn                  = "id"
	boardIdColumn             = "board_id"
	ownerIdColumn             = "owner_id"
	performerIdColumn         = "performer_id"
	contentColumn             = "content"
	statusColumn              = "status"
	completionDateStartColumn = "completed_start"
	completionDateEndColumn   = "completed_end"
	updatedAtColumn           = "updated_at"
	createdAtStartColumn      = "created_at_start"
	createdAtEndColumn        = "created_at_end"
)

func ToNoteFromDesc(note *desc.Note) *model.Note {
	return &model.Note{
		ID:   note.GetId(),
		Info: ToNoteInfoFromDesc(note.GetInfo()),
		UpdatedAt: sql.NullTime{
			Valid: note.GetUpdateAt().IsValid(),
			Time:  note.GetUpdateAt().AsTime(),
		},
		CreatedAt: note.GetCreatedAt().AsTime(),
	}
}

func ToNoteInfoFromDesc(note *desc.NoteInfo) model.NoteInfo {
	isPerformerIdValid := false
	if note.GetPerformerId().GetValue() > 0 {
		isPerformerIdValid = true
	}
	return model.NoteInfo{
		BoardID: note.GetBoardId(),
		OwnerID: note.GetOwnerId(),
		PerformerID: sql.NullInt64{
			Valid: isPerformerIdValid,
			Int64: note.GetPerformerId().GetValue(),
		},
		Content: note.GetContent(),
		Status:  note.GetStatus(),
		CompletionDate: sql.NullTime{
			Valid: note.GetCompletionDate().IsValid(),
			Time:  note.GetCompletionDate().AsTime(),
		},
	}
}

func ToNoteCreateFromDesc(info *desc.NoteCreate) *model.NoteCreate {
	return &model.NoteCreate{
		//BoardID: info.BoardId,
		//OwnerID: info.OwnerId,
		Content: info.Content,
	}
}

func AddOwnerAndBoardIdToNoteCreate(info *model.NoteCreate, ownerID, boardID int64) *model.NoteCreate {
	return &model.NoteCreate{
		Content: info.Content,
		OwnerID: ownerID,
		BoardID: boardID,
	}
}

func ToNoteCompleteFromDesc(req *desc.CompleteRequest) *model.NoteComplete {
	return &model.NoteComplete{
		ID:     req.GetId(),
		Status: req.GetStatus(),
	}
}

func ToNoteUpdateFromDesc(req *desc.UpdateRequest) *model.NoteUpdate {
	return &model.NoteUpdate{
		ID:      req.GetId(),
		Content: req.GetContent(),
	}
}

func ToNoteFromService(note *model.Note) *desc.Note {
	var updateAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updateAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToNoteInfoFromService(note.Info),
		UpdateAt:  updateAt, // Возможно что-то не так будет работать
		CreatedAt: timestamppb.New(note.CreatedAt),
	}
}

func ToNoteInfoFromService(noteInfo model.NoteInfo) *desc.NoteInfo {
	var completionDate *timestamppb.Timestamp
	if noteInfo.CompletionDate.Valid {
		completionDate = timestamppb.New(noteInfo.CompletionDate.Time)
	}

	return &desc.NoteInfo{
		BoardId: noteInfo.BoardID,
		OwnerId: noteInfo.OwnerID,
		PerformerId: &wrapperspb.Int64Value{
			Value: noteInfo.PerformerID.Int64,
		},
		Content:        noteInfo.Content,
		Status:         noteInfo.Status,
		CompletionDate: completionDate,
	}
}

func ToNoteListFromService(notes []*model.Note) []*desc.Note {
	var noteList []*desc.Note
	for _, note := range notes {
		noteList = append(noteList, ToNoteFromService(note))
	}
	return noteList
}

func ToFilters(filters *desc.FilterInfo) map[string]interface{} {
	convertedFilters := make(map[string]interface{})
	//if filters.GetBoardId() != nil {
	//	convertedFilters[boardIdColumn] = filters.GetBoardId()
	//}
	if filters.GetOwnerId() != nil {
		convertedFilters[ownerIdColumn] = filters.GetOwnerId().GetValue()
	}
	if filters.GetPerformerId() != nil {
		convertedFilters[performerIdColumn] = filters.GetPerformerId().GetValue()
	}
	if filters.GetStatus() != nil {
		convertedFilters[statusColumn] = filters.GetStatus().GetValue()
	}
	if filters.GetCreatedAtStart().IsValid() {
		convertedFilters[createdAtStartColumn] = filters.GetCreatedAtStart().AsTime()
	}
	if filters.GetCreatedAtEnd().IsValid() {
		convertedFilters[createdAtEndColumn] = filters.GetCreatedAtEnd().AsTime()
	}
	if filters.GetUpdatedAt().IsValid() {
		convertedFilters[updatedAtColumn] = filters.GetUpdatedAt().AsTime()
	}
	if filters.GetCompletionDateStart().IsValid() {
		convertedFilters[completionDateStartColumn] = filters.GetCompletionDateStart().AsTime()
	}
	if filters.GetCompletionDateEnd().IsValid() {
		convertedFilters[completionDateEndColumn] = filters.GetCompletionDateEnd().AsTime()
	}

	return convertedFilters
}
