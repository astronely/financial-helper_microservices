package wallet

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/financeService/internal/service"
)

var _ def.WalletService = (*serv)(nil)

type serv struct {
	walletRepository repository.WalletRepository
}

func NewService(walletRepository repository.WalletRepository) def.WalletService {
	return &serv{
		walletRepository: walletRepository,
	}
}
