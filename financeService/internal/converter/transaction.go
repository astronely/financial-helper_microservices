package converter

import (
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTransactionInfoFromDesc(req *desc.TransactionInfo) *model.TransactionInfo {
	sum, err := decimal.NewFromString(req.GetSum())
	if err != nil {
		logger.Error("Error converting from balance to decimal",
			"error", err,
			"balance", req.GetSum(),
		)
		sum = decimal.NewFromInt(-1) // TODO: Возможно нужно как-то по другому обрабатывать данный случай
	}

	return &model.TransactionInfo{
		OwnerID:  req.GetOwnerId(),
		WalletID: req.GetWalletId(),
		BoardID:  req.GetBoardId(),
		Sum:      sum,
	}
}

func ToTransactionDetailsInfoFromDesc(req *desc.TransactionDetailsInfo) *model.TransactionDetailsInfo {
	return &model.TransactionDetailsInfo{
		Name:            req.GetName(),
		Category:        req.GetCategory(),
		TransactionDate: req.GetTransactionDate().AsTime(),
	}
}

func ToTransactionFromService(transaction *model.Transaction) *desc.Transaction {
	var updatedAt *timestamppb.Timestamp
	if transaction.UpdatedAt.Valid {
		updatedAt = timestamppb.New(transaction.UpdatedAt.Time)
	}

	return &desc.Transaction{
		Id:        transaction.ID,
		Info:      ToTransactionInfoFromService(transaction.Info),
		DetailsId: transaction.DetailsID,
		UpdatedAt: updatedAt,
		CreatedAt: timestamppb.New(transaction.CreatedAt),
	}
}

func ToTransactionInfoFromService(info model.TransactionInfo) *desc.TransactionInfo {
	return &desc.TransactionInfo{
		OwnerId:  info.OwnerID,
		WalletId: info.WalletID,
		BoardId:  info.BoardID,
		Sum:      info.Sum.String(),
	}
}

func ToTransactionDetailsFromService(details model.TransactionDetails) *desc.TransactionDetails {
	return &desc.TransactionDetails{
		Id:   details.ID,
		Info: ToTransactionDetailsInfoFromService(details.Info),
	}
}

func ToTransactionDetailsInfoFromService(info model.TransactionDetailsInfo) *desc.TransactionDetailsInfo {
	return &desc.TransactionDetailsInfo{
		Name:            info.Name,
		Category:        info.Category,
		TransactionDate: timestamppb.New(info.TransactionDate),
	}
}

func ToTransactionCategoryFromService(info model.TransactionCategory) *desc.TransactionCategory {
	return &desc.TransactionCategory{
		Name:        info.Name,
		Description: info.Description,
	}
}

func ToTransactionListFromService(transactions []*model.Transaction) []*desc.Transaction {
	var transactionList []*desc.Transaction
	for _, transaction := range transactions {
		transactionList = append(transactionList, ToTransactionFromService(transaction))
	}
	return transactionList
}
