package service

import "context"

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context) error
}

type AccessService interface {
	Check(ctx context.Context, endpointAddress string) (bool, error)
}
