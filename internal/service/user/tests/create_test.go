package tests

import (
	"context"
	"fmt"
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	"github.com/astronely/financial-helper_microservices/internal/repository/mocks"
	"github.com/astronely/financial-helper_microservices/internal/service/user"
	"github.com/astronely/financial-helper_microservices/internal/utils"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_serv_Create(t *testing.T) {
	type userRepositoryMockFunc func() repository.UserRepository

	type args struct {
		ctx      context.Context
		info     *model.UserInfo
		password string
	}

	var (
		ctx = context.Background()

		email    = gofakeit.Email()
		name     = gofakeit.Name()
		token, _ = utils.GenerateToken(model.UserInfo{
			Name:  name,
			Email: email,
		}, []byte("testSecretKey"), 360)
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
	}{
		{
			name: "success case",
			args: args{
				ctx:      context.Background(),
				info:     info,
				password: password,
			},
			err:   nil,
			want:  1,
			want1: token,
			userRepositoryMock: func() repository.UserRepository {
				mock := mocks.NewUserRepository(t)
				mock.On("Create", ctx, info, password).Return(int64(1), nil)
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
				mock.On("Create", ctx, info, password).Return(int64(0), repoErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepositoryMock := tt.userRepositoryMock()
			defer func() {
				if mock, ok := userRepositoryMock.(*mocks.UserRepository); ok {
					mock.AssertExpectations(t)
				}
			}()

			service := user.NewMockService(userRepositoryMock)

			resID, resToken, err := service.Create(tt.args.ctx, tt.args.info, tt.args.password)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resID)
			require.Equal(t, tt.want1, resToken)
		})
	}
}
