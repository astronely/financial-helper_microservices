package transaction

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

type Implementation struct {
	desc.UnimplementedTransactionV1Server
	service service.TransactionService
}

func NewImplementation(service service.TransactionService) *Implementation {
	return &Implementation{
		service: service,
	}
}
