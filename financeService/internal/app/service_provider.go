package app

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/closer"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	apiTransaction "github.com/astronely/financial-helper_microservices/financeService/internal/api/transaction"
	"github.com/astronely/financial-helper_microservices/financeService/internal/api/wallet"
	"github.com/astronely/financial-helper_microservices/financeService/internal/config"
	"github.com/astronely/financial-helper_microservices/financeService/internal/config/env"
	"github.com/astronely/financial-helper_microservices/financeService/internal/repository"
	walletRepository "github.com/astronely/financial-helper_microservices/financeService/internal/repository/wallet"
	"github.com/astronely/financial-helper_microservices/financeService/internal/service"
	walletService "github.com/astronely/financial-helper_microservices/financeService/internal/service/wallet"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db/pg"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db/transaction"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	pgConfig   config.PGConfig

	dbClient  db.Client
	txManager db.TxManager

	// TODO: Services, Repositories and Implementations
	walletService service.WalletService

	walletRepository repository.WalletRepository

	walletImpl      *wallet.Implementation
	transactionImpl *apiTransaction.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			panic("Cannot load GRPC Config" + err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			panic("Cannot load PG Config" + err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			panic("Error initializing DB client/CREATE" + err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			panic("Error initializing DB client/PING")
		}
		closer.Add(func() error {
			err := cl.Close()
			if err != nil {
				return err
			}
			logger.Info("Successfully closed DB client")
			return nil
		})

		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) WalletService(ctx context.Context) service.WalletService {
	if s.walletService == nil {
		s.walletService = walletService.NewService(s.WalletRepository(ctx), s.TxManager(ctx))
	}
	return s.walletService
}

func (s *serviceProvider) WalletRepository(ctx context.Context) repository.WalletRepository {
	if s.walletRepository == nil {
		s.walletRepository = walletRepository.NewRepository(s.DBClient(ctx))
	}
	return s.walletRepository
}

func (s *serviceProvider) WalletImpl(ctx context.Context) *wallet.Implementation {
	if s.walletImpl == nil {
		s.walletImpl = wallet.NewImplementation(s.WalletService(ctx))
	}
	return s.walletImpl
}

func (s *serviceProvider) TransactionImpl(ctx context.Context) *apiTransaction.Implementation {
	if s.transactionImpl == nil {
		s.transactionImpl = apiTransaction.NewImplementation()
	}
	return s.transactionImpl
}
