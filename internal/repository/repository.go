package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo, password string) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context, limit uint64, offset uint64) ([]*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, info *model.UpdateUserInfo) (int64, error)
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error)
}

type AccessRepository interface {
	CheckUser(ctx context.Context, id int64) (bool, error)
}
