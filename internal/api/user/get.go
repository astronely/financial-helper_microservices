package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		log.Printf("error getting user: %v", err)
		return nil, err
	}
	log.Printf("Get User: %v", userObj.ID)
	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}
