package access

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	isAllowed, err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &desc.CheckResponse{
		IsAllowed: isAllowed,
	}, nil
}
