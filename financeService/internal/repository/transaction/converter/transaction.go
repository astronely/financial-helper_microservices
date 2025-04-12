package converter

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/model"
)

func ToTransactionFromRepo(transaction *modelRepo.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:                 transaction.ID,
		Info:               ToTransactionInfoFromRepo(transaction.Info),
		CreatedAt:          transaction.CreatedAt,
		UpdatedAt:          transaction.UpdatedAt,
		TransactionDetails: ToTransactionDetailsFromRepo(transaction.TransactionDetails),
	}
}

func ToTransactionInfoFromRepo(transactionInfo modelRepo.TransactionInfo) model.TransactionInfo {
	return model.TransactionInfo{
		OwnerID:  transactionInfo.OwnerID,
		BoardID:  transactionInfo.BoardID,
		WalletID: transactionInfo.WalletID,
		Sum:      transactionInfo.Sum,
	}
}

func ToTransactionDetailsFromRepo(transactionDetails modelRepo.TransactionDetails) model.TransactionDetails {
	return model.TransactionDetails{
		ID:                  transactionDetails.ID,
		Info:                ToTransactionDetailsInfoFromRepo(transactionDetails),
		TransactionCategory: ToTransactionCategoryFromRepo(transactionDetails.TransactionCategory),
	}
}

func ToTransactionDetailsInfoFromRepo(transactionDetails modelRepo.TransactionDetails) model.TransactionDetailsInfo {
	return model.TransactionDetailsInfo{
		Name:            transactionDetails.Name,
		TransactionDate: transactionDetails.TransactionDate,
	}
}

func ToTransactionCategoryFromRepo(transactionCategory modelRepo.TransactionCategory) model.TransactionCategory {
	return model.TransactionCategory{
		ID:          transactionCategory.ID,
		Name:        transactionCategory.Name,
		Description: transactionCategory.Description,
	}
}

func ToTransactionListFromRepo(transactions []*modelRepo.Transaction) []*model.Transaction {
	var transactionList []*model.Transaction
	for _, transaction := range transactions {
		transactionList = append(transactionList, ToTransactionFromRepo(transaction))
	}
	return transactionList
}
