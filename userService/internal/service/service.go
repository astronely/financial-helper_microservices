package service

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo, password string) (int64, string, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context, limit uint64, offset uint64) ([]*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, info *model.UpdateUserInfo) (int64, error)
}

//type AuthService interface {
//	Login(ctx context.Context, email string, password string) (string, error)
//	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
//	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
//}
//
//type AccessService interface {
//	Check(ctx context.Context, endpointAddress string) (bool, error)
//}
