package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo, password string) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error)
}
