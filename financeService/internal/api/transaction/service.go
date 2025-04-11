package transaction

import (
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
)

type Implementation struct {
	desc.UnimplementedTransactionV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
