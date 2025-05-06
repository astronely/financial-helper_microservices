package converter

import (
	"database/sql"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToBoardFromService(info *model.Board) *desc.Board {
	var updatedAt *timestamppb.Timestamp
	if info.UpdatedAt.Valid {
		updatedAt = timestamppb.New(info.UpdatedAt.Time)
	}

	return &desc.Board{
		Id:        info.ID,
		Info:      ToBoardInfoFromService(info.Info),
		UpdatedAt: updatedAt,
		CreatedAt: timestamppb.New(info.CreatedAt),
	}
}

func ToBoardCreateFromDesc(req *desc.CreateRequest) *model.BoardCreate {
	return &model.BoardCreate{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
}

func ToBoardInfoFromDesc(req *desc.BoardInfo) *model.BoardInfo {
	return &model.BoardInfo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		OwnerID:     req.GetOwnerId(),
	}
}

func ToBoardInfoFromService(info model.BoardInfo) *desc.BoardInfo {
	return &desc.BoardInfo{
		Name:        info.Name,
		Description: info.Description,
		OwnerId:     info.OwnerID,
	}
}

func ToBoardUsersFromService(info []*model.BoardUser) []*desc.BoardUser {
	var users []*desc.BoardUser
	for _, user := range info {
		users = append(users, ToBoardUserFromService(user))
	}
	return users
}

func ToBoardUserFromService(user *model.BoardUser) *desc.BoardUser {
	return &desc.BoardUser{
		BoardId: user.BoardID,
		UserId:  user.UserID,
		Role:    user.Role,
	}
}

func ToBoardListFromService(boards []*model.Board) []*desc.Board {
	var boardList []*desc.Board
	for _, board := range boards {
		boardList = append(boardList, ToBoardFromService(board))
	}
	return boardList
}

func ToGenerateInviteFromDesc(req *desc.GenerateInviteRequest) *model.GenerateInviteInfo {
	return &model.GenerateInviteInfo{
		BoardID: req.GetBoardId(),
		//UserID:  req.GetUserId(),
		Role: req.GetRole(),
	}
}

func ToBoardUpdateFromDesc(req *desc.UpdateRequest) *model.BoardUpdate {
	var name, description sql.NullString

	if req.GetName() != nil {
		name = sql.NullString{
			String: req.GetName().GetValue(),
			Valid:  true,
		}
	}
	if req.GetDescription() != nil {
		description = sql.NullString{
			String: req.GetDescription().GetValue(),
			Valid:  true,
		}
	}

	return &model.BoardUpdate{
		ID:          req.GetId(),
		Name:        name,
		Description: description,
	}
}

func ToJoinInfoFromDesc(req *desc.JoinRequest) *model.JoinInfo {
	return &model.JoinInfo{
		Token: req.GetToken(),
		//ID:    req.GetId(),
	}
}

func ToGenerateInvite(info *model.GenerateInviteInfo) model.GenerateInviteInfo {
	return model.GenerateInviteInfo{
		BoardID: info.BoardID,
		//UserID:  info.UserID,
		Role: info.Role,
	}
}
