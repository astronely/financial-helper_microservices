package wallet

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
)

type Implementation struct {
	desc.UnimplementedWalletV1Server
	service service.WalletService
}

func NewImplementation(service service.WalletService) *Implementation {
	return &Implementation{
		service: service,
	}
}
