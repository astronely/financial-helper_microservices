package tests

import (
	"context"
	"fmt"
	"github.com/astronely/financial-helper_microservices/userService/internal/api/user"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/internal/service/mocks"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImplementation_Create(t *testing.T) {
	type userServiceMockFunc func() service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 10)
		token    = "token"

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Email: email,
				Name:  name,
			},
			Password: password,
		}

		info = &model.UserInfo{
			Email: email,
			Name:  name,
		}

		res = &desc.CreateResponse{
			Id:    id,
			Token: token,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mockService.On("Create", ctx, info, password).Return(id, token, nil)
				return mockService
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func() service.UserService {
				mockService := mocks.NewUserService(t)
				mockService.On("Create", ctx, info, password).Return(int64(0), token, serviceErr)
				return mockService
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock()
			defer func() {
				if mock, ok := userServiceMock.(*mocks.UserService); ok {
					mock.AssertExpectations(t)
				}
			}()

			api := user.NewImplementation(userServiceMock)

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
