package app

import (
	"context"
	"github.com/astronely/financial-helper_microservices/internal/api/access"
	"github.com/astronely/financial-helper_microservices/internal/api/auth"
	"github.com/astronely/financial-helper_microservices/internal/api/user"
	"github.com/astronely/financial-helper_microservices/internal/config"
	"github.com/astronely/financial-helper_microservices/internal/config/env"
	"github.com/astronely/financial-helper_microservices/internal/repository"
	authRepository "github.com/astronely/financial-helper_microservices/internal/repository/auth"
	userRepository "github.com/astronely/financial-helper_microservices/internal/repository/user"
	"github.com/astronely/financial-helper_microservices/internal/service"
	accessService "github.com/astronely/financial-helper_microservices/internal/service/access"
	authService "github.com/astronely/financial-helper_microservices/internal/service/auth"
	userService "github.com/astronely/financial-helper_microservices/internal/service/user"
	"github.com/astronely/financial-helper_microservices/pkg/client/db"
	"github.com/astronely/financial-helper_microservices/pkg/client/db/pg"
	"github.com/astronely/financial-helper_microservices/pkg/client/db/transaction"
	"github.com/astronely/financial-helper_microservices/pkg/closer"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	authRepository repository.AuthRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			panic("Error loading PG config")
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			panic("Error loading GRPC config")
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			panic("Error loading HTTP config")
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			panic("Error loading Swagger config")
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			panic("Error initializing DB client")
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

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(s.AuthService(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
