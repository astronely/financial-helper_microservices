package transaction

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/financeService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

var _ def.TransactionService = (*serv)(nil)

type serv struct {
	transactionRepository repository.TransactionRepository
	txManager             db.TxManager
}

func NewService(transactionRepository repository.TransactionRepository, txManager db.TxManager) def.TransactionService {
	return &serv{
		transactionRepository: transactionRepository,
		txManager:             txManager,
	}
}
