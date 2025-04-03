package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/internal/converter"
	"github.com/astronely/financial-helper_microservices/userService/pkg/logger"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	usersObj, err := i.userService.List(ctx, uint64(req.GetLimit()), uint64(req.GetOffset()))
	if err != nil {
		logger.Error("Error getting user",
			"err", err.Error(),
		)
		return nil, err
	}
	logger.Debug("Get List of Users",
		"users number", len(usersObj),
	)

	return &desc.ListResponse{
		Users: converter.ToUserListFromService(usersObj),
	}, nil
}
