package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
