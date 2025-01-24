package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, token, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()), req.GetPassword())
	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	log.Printf("Created user with id: %v", id)

	return &desc.CreateResponse{
		Id:    id,
		Token: token,
	}, nil
}
