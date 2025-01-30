package service

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
