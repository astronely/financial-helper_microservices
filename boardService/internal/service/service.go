package service

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

type BoardService interface {
	Create(ctx context.Context, info *model.BoardCreate) (int64, error)
	CreateUser(ctx context.Context, info *model.BoardUserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.Board, error)
	GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error)
	ListByUserId(ctx context.Context) ([]*model.Board, error)
	ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error)
	Update(ctx context.Context, info *model.BoardUpdate) (int64, error)
	Delete(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, userID int64) error
	CompareUserAndBoard(ctx context.Context) (bool, error)
	CompareUserAndBoardRaw(ctx context.Context, userId int64, boardId int64) (bool, error)
	CheckUserInBoardWithContext(ctx context.Context, boardID int64) error
	SetBoard(ctx context.Context, boardID int64) error

	JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error)
	GenerateInvite(ctx context.Context) (string, error)
}
