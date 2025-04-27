package board

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository"
	modelRepo "github.com/astronely/financial-helper_microservices/boardService/internal/repository/redis/board/model"
	"github.com/astronely/financial-helper_microservices/boardService/pkg/client/cache"
	"github.com/google/uuid"
	"time"
)

const (
	inviteKey = "invites"

	boardIdName = "boardId"
)

type repo struct {
	client cache.RedisClient
}

func NewRepository(client cache.RedisClient) repository.BoardRedisRepository {
	return &repo{client: client}
}

//func (r *repo) Create(ctx context.Context, info *model.BoardInfo) (int64, error) {
//	return -1, errors.New("not implemented")
//}
//
//func (r *repo) CreateUser(ctx context.Context, info *model.BoardUser) (int64, error) {
//	return -1, errors.New("not implemented")
//}
//
//func (r *repo) Get(ctx context.Context, id int64) (*model.Board, error) {
//	return nil, errors.New("not implemented")
//}
//
//func (r *repo) GetUsers(ctx context.Context, boardId int64) ([]*model.BoardUser, error) {
//	return nil, errors.New("not implemented")
//}
//
//func (r *repo) ListByUserId(ctx context.Context, userId int64) ([]*model.Board, error) {
//	return nil, errors.New("not implemented")
//}
//
//func (r *repo) ListByOwnerId(ctx context.Context, ownerId int64) ([]*model.Board, error) {
//	return nil, errors.New("not implemented")
//}
//
//func (r *repo) Update(ctx context.Context, info *model.BoardUpdate) (int64, error) {
//	return -1, errors.New("not implemented")
//}
//
//func (r *repo) Delete(ctx context.Context, id int64) error {
//	return errors.New("not implemented")
//}

func (r *repo) GenerateInvite(ctx context.Context, info *model.GenerateInviteInfo) (string, error) {
	token := uuid.New().String()
	key := inviteKey + token
	converted := modelRepo.InvitationData{
		BoardID: info.BoarID,
		UserID:  info.UserID,
		Role:    info.Role,
	}

	bytesData, err := json.Marshal(converted)
	if err != nil {
		return "", err
	}

	err = r.client.Set(ctx, key, bytesData)
	if err != nil {
		logger.Error("error generate invite | Redis",
			"error", err.Error(),
		)
		return "", err
	}

	err = r.client.Expire(ctx, key, time.Minute*15)
	if err != nil {
		logger.Error("error expire | Redis",
			"error", err.Error(),
		)
		return "", err
	}

	return token, nil
}

func (r *repo) JoinBoard(ctx context.Context, info *model.JoinInfo) (*model.GenerateInviteInfo, error) {
	key := inviteKey + info.Token

	inviteInfo, err := r.client.Get(ctx, key)
	if err != nil {
		logger.Error("error join board | Redis",
			"error", err.Error(),
		)
		return nil, err
	}

	var convertedInfo *model.GenerateInviteInfo
	err = json.Unmarshal(inviteInfo.([]byte), &convertedInfo)
	if err != nil {
		return nil, errors.New("invalid invite info")
	}
	return convertedInfo, nil
	//switch convertedInfo.(type) {
	//case model.GenerateInviteInfo:
	//	return inviteInfo.(*model.GenerateInviteInfo), nil
	//default:
	//	return nil, errors.New("invalid invite info")
	//}
}
