package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/astronely/financial-helper_microservices/userService/internal/api/user"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/internal/service/mocks"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestImplementation_Get(t *testing.T) {
	type userServiceMockFunc func() service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.User{
			ID: id,
			Info: model.UserInfo{
				Email: email,
				Name:  name,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Email: email,
					Name:  name,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func() service.UserService {
				mockService := mocks.NewUserService(t)
				mockService.On("Get", ctx, id).Return(serviceRes, nil)
				return mockService
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func() service.UserService {
				mockService := mocks.NewUserService(t)
				mockService.On("Get", ctx, id).Return(nil, serviceErr)
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			userServiceMock := tt.userServiceMock()
			defer func() {
				if mock, ok := userServiceMock.(*mocks.UserService); ok {
					mock.AssertExpectations(t)
				}
			}()

			api := user.NewImplementation(userServiceMock)

			res, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)

		})
	}
}
