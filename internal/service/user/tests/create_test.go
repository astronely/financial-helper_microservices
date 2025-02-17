package tests

import (
	"context"
	"fmt"
	"github.com/astronely/financial-helper_microservices/internal/config"
	tokenMocks "github.com/astronely/financial-helper_microservices/internal/config/mocks"
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/repository/mocks"
	"github.com/astronely/financial-helper_microservices/internal/service/user"
	"github.com/astronely/financial-helper_microservices/internal/utils"
	"github.com/brianvoe/gofakeit/v7"
	mockLib "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_serv_Create(t *testing.T) {
	type userRepositoryMockFunc func() repository.UserRepository
	type tokenConfigMockFunc func() config.TokenConfig

	type args struct {
		ctx      context.Context
		info     *model.UserInfo
		password string
	}

	var (
		ctx      = context.Background()
		id       = gofakeit.Int64()
		email    = gofakeit.Email()
		name     = gofakeit.Name()
		token, _ = utils.GenerateToken(id,
			model.UserInfo{
				Name:  name,
				Email: email,
			}, []byte("testSecretKey"), 360*time.Second)
		password = gofakeit.Password(true, true, true, true, false, 10)
		info     = &model.UserInfo{
			Email: email,
			Name:  name,
		}

		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		want1              string
		err                error
		userRepositoryMock userRepositoryMockFunc
		tokenConfigMock    tokenConfigMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:      context.Background(),
				info:     info,
				password: password,
			},
			err:   nil,
			want:  id,
			want1: token,
			userRepositoryMock: func() repository.UserRepository {
				mock := mocks.NewUserRepository(t)
				mock.On("Create", ctx, info, password).Return(id, nil)
				return mock
			},
			tokenConfigMock: func() config.TokenConfig {
				mock := tokenMocks.NewTokenConfig(t)
				mock.On("RefreshTokenKey").Return("testSecretKey")
				mock.On("RefreshTokenExpirationTime").Return(time.Duration(360 * time.Second))
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:      context.Background(),
				info:     info,
				password: password,
			},
			err:   repoErr,
			want:  0,
			want1: "",
			userRepositoryMock: func() repository.UserRepository {
				mock := mocks.NewUserRepository(t)
				mock.On("Create", mockLib.Anything, info, password).Return(int64(0), repoErr)
				return mock
			},
			tokenConfigMock: func() config.TokenConfig {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepositoryMock := tt.userRepositoryMock()
			tokenConfigMock := tt.tokenConfigMock()

			defer func() {
				if mock, ok := userRepositoryMock.(*mocks.UserRepository); ok {
					mock.AssertExpectations(t)
				}
				if tokenMock, ok := tokenConfigMock.(*tokenMocks.TokenConfig); ok {
					tokenMock.AssertExpectations(t)
				}
			}()

			service := user.NewMockService(userRepositoryMock, tokenConfigMock)

			resID, resToken, err := service.Create(tt.args.ctx, tt.args.info, tt.args.password)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resID, "Id dont match")
			require.Equal(t, tt.want1, resToken, "Token dont match")
		})
	}
}
