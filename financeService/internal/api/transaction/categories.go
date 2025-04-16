package transaction

import (
	"context"
	"github.com/astronely/financial-helper_microservices/financeService/internal/converter"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Categories(ctx context.Context, req *emptypb.Empty) (*desc.CategoriesResponse, error) {
	categories, err := i.service.Categories(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.CategoriesResponse{
		Categories: converter.ToCategoriesFromService(categories),
	}, nil
}
