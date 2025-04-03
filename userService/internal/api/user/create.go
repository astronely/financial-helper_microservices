package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/internal/converter"
	"github.com/astronely/financial-helper_microservices/userService/pkg/logger"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, token, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()), req.GetPassword())
	if err != nil {
		logger.Error("Error creating user",
			"err", err.Error(),
		)
		return nil, err
	}

	logger.Debug("Created user",
		"id", id,
	)

	return &desc.CreateResponse{
		Id:    id,
		Token: token,
	}, nil
}
