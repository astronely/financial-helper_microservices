package wallet

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/financeService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

var _ def.WalletService = (*serv)(nil)

type serv struct {
	walletRepository repository.WalletRepository
	txManager        db.TxManager
}

func NewService(walletRepository repository.WalletRepository, txManager db.TxManager) def.WalletService {
	return &serv{
		walletRepository: walletRepository,
		txManager:        txManager,
	}
}
