package repository

import (
	"context"
	"github.com/astronely/financial-helper_microservices/authService/internal/model"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error)
}

type AccessRepository interface {
	CheckUser(ctx context.Context, id int64) (bool, error)
}
