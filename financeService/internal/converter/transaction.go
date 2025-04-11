package converter

import (
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"github.com/shopspring/decimal"
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
