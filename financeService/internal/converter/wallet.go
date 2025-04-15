package converter

import (
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToWalletFromService(wallet *model.Wallet) *desc.Wallet {
	//var updatedAt *timestamppb.Timestamp
	//if wallet.UpdatedAt.Valid {
	//	updatedAt = timestamppb.New(wallet.UpdatedAt.Time)
	//}

	return &desc.Wallet{
		Id:        wallet.ID,
		Info:      ToWalletInfoFromService(wallet.Info),
		CreatedAt: timestamppb.New(wallet.CreatedAt),
		// TODO: UpdatedAt
	}
}

func ToWalletListFromService(wallets []*model.Wallet) []*desc.Wallet {
	var walletList []*desc.Wallet
	for _, wallet := range wallets {
		walletList = append(walletList, ToWalletFromService(wallet))
	}
	return walletList
}

func ToWalletInfoFromService(walletInfo model.WalletInfo) *desc.WalletInfo {
	return &desc.WalletInfo{
		OwnerId: walletInfo.OwnerID,
		BoardId: walletInfo.BoardID,
		Name:    walletInfo.Name,
		Balance: walletInfo.Balance.String(),
	}
}

func ToUpdateWalletInfoFromDesc(wallet *desc.UpdateRequest) *model.WalletUpdateInfo {
	balance, err := decimal.NewFromString(wallet.GetInfo().GetBalance().GetValue())
	if err != nil {
		logger.Error("Error converting from balance to decimal",
			"error", err,
			"balance", wallet.GetInfo().GetBalance().GetValue(),
		)
		balance = decimal.NewFromInt(-1) // TODO: Возможно нужно как-то по другому обрабатывать данный случай
	}

	return &model.WalletUpdateInfo{
		ID:      wallet.GetId(),
		Name:    wallet.GetInfo().GetName().GetValue(),
		Balance: balance,
	}
}

func ToWalletInfoFromDesc(walletInfo *desc.WalletInfo) *model.WalletInfo {
	balance, err := decimal.NewFromString(walletInfo.GetBalance())
	if err != nil {
		logger.Error("Error converting from balance to decimal",
			"error", err,
			"balance", walletInfo.GetBalance(),
		)
		balance = decimal.NewFromInt(-1) // TODO: Возможно нужно как-то по другому обрабатывать данный случай
	}

	return &model.WalletInfo{
		OwnerID: walletInfo.GetOwnerId(),
		BoardID: walletInfo.GetBoardId(),
		Name:    walletInfo.GetName(),
		Balance: balance,
	}
}

func ToWalletFromDesc(wallet *desc.Wallet) *model.Wallet {
	return &model.Wallet{
		ID:   wallet.GetId(),
		Info: *ToWalletInfoFromDesc(wallet.GetInfo()),
	}
}
