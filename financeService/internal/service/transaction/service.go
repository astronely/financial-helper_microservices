package transaction

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/config"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/financeService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

var _ def.TransactionService = (*serv)(nil)

type serv struct {
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
	txManager             db.TxManager

	tokenConfig config.TokenConfig
}

func NewService(transactionRepository repository.TransactionRepository,
	walletRepository repository.WalletRepository, txManager db.TxManager,
	tokenConfig config.TokenConfig) def.TransactionService {
	return &serv{
		transactionRepository: transactionRepository,
		walletRepository:      walletRepository,
		txManager:             txManager,
		tokenConfig:           tokenConfig,
	}
}
