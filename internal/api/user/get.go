package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/converter"
	"github.com/astronely/financial-helper_microservices/internal/logger"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("Error getting user",
			"err", err,
		)
		return nil, err
	}
	logger.Debug("Get User",
		"id", userObj.ID,
	)

	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}
