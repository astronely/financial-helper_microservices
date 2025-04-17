package converter

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/transaction/model"
)

func ToTransactionFromRepo(transaction *modelRepo.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:                 transaction.ID,
		Info:               ToTransactionInfoFromRepo(transaction.Info),
		DetailsId:          transaction.DetailsID,
		CreatedAt:          transaction.CreatedAt,
		UpdatedAt:          transaction.UpdatedAt,
		TransactionDetails: ToTransactionDetailsFromRepo(transaction.TransactionDetails),
	}
}

func ToTransactionInfoFromRepo(transactionInfo modelRepo.TransactionInfo) model.TransactionInfo {
	return model.TransactionInfo{
		OwnerID:      transactionInfo.OwnerID,
		BoardID:      transactionInfo.BoardID,
		FromWalletID: transactionInfo.FromWalletID,
		ToWalletID:   transactionInfo.ToWalletID,
		Amount:       transactionInfo.Amount,
		Type:         transactionInfo.Type,
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
		Category:        transactionDetails.Category,
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
