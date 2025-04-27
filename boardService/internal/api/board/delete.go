package board

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
