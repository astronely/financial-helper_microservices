package converter

import (
	"database/sql"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

const (
	categoryColumn           = "category"
	transactionDateColumn    = "transaction_date"
	transactionDateColumnEnd = transactionDateColumn + "_end"

	ownerIdColumn      = "owner_id"
	fromWalletIdColumn = "from_wallet_id"
	toWalletIdColumn   = "to_wallet_id"
	typeColumn         = "type"
	nameColumn         = "name"
)

func ToTransactionInfoFromDesc(req *desc.CreateTransactionInfo) *model.TransactionInfo {
	amount, err := decimal.NewFromString(req.GetAmount())
	if err != nil {
		logger.Error("Error converting from balance to decimal",
			"error", err,
			"balance", req.GetAmount(),
		)
		amount = decimal.NewFromInt(-1) // TODO: Возможно нужно как-то по другому обрабатывать данный случай
	}
	var isValidToWalletId = false

	if req.GetToWalletId().GetValue() > 0 {
		isValidToWalletId = true
	}

	return &model.TransactionInfo{
		//OwnerID:      req.GetOwnerId(),
		FromWalletID: req.GetFromWalletId(),
		ToWalletID: sql.NullInt64{
			Valid: isValidToWalletId,
			Int64: req.GetToWalletId().GetValue(),
		},
		//BoardID: req.GetBoardId(),
		Amount: amount,
		Type:   req.GetType(),
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
		DetailsId: transaction.DetailsId,
		UpdatedAt: updatedAt,
		CreatedAt: timestamppb.New(transaction.CreatedAt),
	}
}

func ToTransactionInfoFromService(info model.TransactionInfo) *desc.TransactionInfo {
	returningDesc := &desc.TransactionInfo{
		OwnerId:      info.OwnerID,
		FromWalletId: info.FromWalletID,
		BoardId:      info.BoardID,
		Amount:       info.Amount.String(),
		Type:         info.Type,
	}

	if info.ToWalletID.Valid {
		returningDesc.ToWalletId = &wrapperspb.Int64Value{Value: info.ToWalletID.Int64}
	}

	return returningDesc
}

func AddOwnerAndBoardIdToTransactionInfo(transactionInfo *model.TransactionInfo, ownerID, boardID int64) *model.TransactionInfo {
	return &model.TransactionInfo{
		OwnerID:      ownerID,
		FromWalletID: transactionInfo.FromWalletID,
		ToWalletID:   transactionInfo.ToWalletID,
		BoardID:      boardID,
		Amount:       transactionInfo.Amount,
		Type:         transactionInfo.Type,
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
		Id:          info.ID,
		Name:        info.Name,
		Description: info.Description,
	}
}

func ToTransactionListFromService(transactions []*model.Transaction) []*desc.GetResponse {
	var transactionList []*desc.GetResponse
	for _, transaction := range transactions {
		transactionList = append(transactionList, &desc.GetResponse{
			Transaction: ToTransactionFromService(transaction),
			Details:     ToTransactionDetailsFromService(transaction.TransactionDetails),
			Category:    ToTransactionCategoryFromService(transaction.TransactionDetails.TransactionCategory),
		})
	}
	return transactionList
}

func ToTransactionInfoUpdateFromDesc(req *desc.UpdateRequest) *model.TransactionInfoUpdate {
	amount, err := decimal.NewFromString(req.GetInfo().GetAmount().GetValue())
	if err != nil {
		logger.Error("Error converting from balance to decimal",
			"error", err,
			"balance", req.GetInfo().GetAmount().GetValue(),
		)
		amount = decimal.NewFromInt(-1) // TODO: Возможно нужно как-то по другому обрабатывать данный случай
	}
	return &model.TransactionInfoUpdate{
		ID:           req.GetId(),
		FromWalletID: req.GetInfo().GetFromWalletId().GetValue(),
		ToWalletID:   req.GetInfo().GetToWalletId().GetValue(),
		Amount:       amount,
		Type:         req.GetInfo().GetType().GetValue(),
	}
}

func ToTransactionDetailsUpdateFromDesc(req *desc.UpdateRequest) *model.TransactionDetailsUpdate {
	transactionDetailsUpdate := &model.TransactionDetailsUpdate{
		ID:              req.GetId(),
		Name:            req.GetInfo().GetName().GetValue(),
		Category:        req.GetInfo().GetCategory().GetValue(),
		TransactionDate: time.Time{},
	}

	if req.GetInfo().GetTransactionDate().IsValid() {
		transactionDetailsUpdate.TransactionDate = req.GetInfo().GetTransactionDate().AsTime()
	}
	return transactionDetailsUpdate
}

func Filters(req *desc.FilterInfo) map[string]interface{} {
	filters := map[string]interface{}{}
	if req.GetCategory() != nil {
		filters[categoryColumn] = req.GetCategory().GetValue()
	}
	if req.GetTransactionDate().IsValid() {
		filters[transactionDateColumn] = req.GetTransactionDate().AsTime()
	}
	if req.GetTransactionDateEnd().IsValid() {
		filters[transactionDateColumnEnd] = req.GetTransactionDateEnd().AsTime()
	}
	if req.GetOwnerId() != nil {
		filters[ownerIdColumn] = req.GetOwnerId().GetValue()
	}
	if req.GetFromWalletId() != nil {
		filters[fromWalletIdColumn] = req.GetFromWalletId().GetValue()
	}
	if req.GetToWalletId() != nil {
		filters[toWalletIdColumn] = req.GetToWalletId().GetValue()
	}
	if req.GetType() != nil {
		filters[typeColumn] = req.GetType().GetValue()
	}
	if req.GetName() != nil {
		filters[nameColumn] = req.GetName().GetValue()
	}

	return filters
}

func ToCategoriesFromService(categories []*model.TransactionCategory) []*desc.TransactionCategory {
	var descCategories []*desc.TransactionCategory
	for _, category := range categories {
		descCategories = append(descCategories, ToTransactionCategoryFromService(*category))
	}
	return descCategories
}
