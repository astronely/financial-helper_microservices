package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
)

type BoardRepository interface {
	Create(ctx context.Context, info *model.BoardInfo) (int64, error)
	CreateUser(ctx context.Context, info *model.BoardUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.Board, error)
	GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error)
	ListByUserId(ctx context.Context, userId int64) ([]*model.Board, error)
	ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error)
	Update(ctx context.Context, info *model.BoardUpdate) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type BoardRedisRepository interface {
	JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error)
	GenerateInvite(ctx context.Context, info *model.GenerateInviteInfo) (string, error)
}
