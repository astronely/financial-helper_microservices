package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/converter"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
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
