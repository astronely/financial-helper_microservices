package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/repository"
	"github.com/astronely/financial-helper_microservices/userService/internal/repository/mocks"
	user2 "github.com/astronely/financial-helper_microservices/userService/internal/service/user"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_serv_Get(t *testing.T) {

	type userRepositoryMockFunc func() repository.UserRepository

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx       = context.Background()
		id        = gofakeit.Int64()
		email     = gofakeit.Email()
		name      = gofakeit.Name()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		info = model.UserInfo{
			Email: email,
			Name:  name,
		}

		user = &model.User{
			ID:        id,
			Info:      info,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: user,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				mock := mocks.NewUserRepository(t)
				mock.On("Get", ctx, id).Return(user, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func() repository.UserRepository {
				mock := mocks.NewUserRepository(t)
				mock.On("Get", ctx, id).Return(nil, repoErr)
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

			service := user2.NewMockService(userRepositoryMock)

			res, err := service.Get(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
