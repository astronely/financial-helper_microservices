package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/userService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.userService.Update(ctx, converter.ToUpdateUserInfoFromDesc(req))

	if err != nil {
		logger.Error("Error updating user",
			"id", req.GetId(),
		)
		return nil, err
	}
	logger.Debug("Successfully updated user",
		"id", id,
	)
	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
