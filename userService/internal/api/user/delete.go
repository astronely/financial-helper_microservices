package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/pkg/logger"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		logger.Error("Error deleting user",
			"id", req.Id,
		)
		return &emptypb.Empty{}, err
	}

	logger.Debug("Successfully deleted user",
		"id", req.Id,
	)

	return &emptypb.Empty{}, nil
}
