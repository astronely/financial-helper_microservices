package converter

import (
	"github.com/astronely/financial-helper_microservices/financeService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/financeService/internal/repository/wallet/model"
)

func ToWalletFromRepo(wallet *modelRepo.Wallet) *model.Wallet {
	return &model.Wallet{
		ID:        wallet.ID,
		Info:      ToWalletInfoFromRepo(wallet.Info),
		CreatedAt: wallet.CreatedAt,
		// UpdatedAt: // TODO
	}
}

func ToWalletsFromRepo(wallets []*modelRepo.Wallet) []*model.Wallet {
	var walletsFromRepo []*model.Wallet
	for _, wallet := range wallets {
		walletsFromRepo = append(walletsFromRepo, ToWalletFromRepo(wallet))
	}
	return walletsFromRepo
}

func ToWalletInfoFromRepo(walletInfo modelRepo.WalletInfo) model.WalletInfo {
	return model.WalletInfo{
		OwnerID: walletInfo.OwnerID,
		BoardID: walletInfo.BoardID,
		Name:    walletInfo.Name,
		Balance: walletInfo.Balance,
	}
}
