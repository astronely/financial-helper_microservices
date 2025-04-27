package board

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func (i *Implementation) ListByOwnerId(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	return nil, nil
}
