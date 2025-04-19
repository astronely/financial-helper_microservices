package transaction

import (
	"context"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.service.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
