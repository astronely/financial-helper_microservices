package app

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/closer"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/boardService/internal/api/board"
	"github.com/astronely/financial-helper_microservices/boardService/internal/config"
	"github.com/astronely/financial-helper_microservices/boardService/internal/config/env"
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository"
	boardRepository "github.com/astronely/financial-helper_microservices/boardService/internal/repository/pg/board"
	boardRedisRepository "github.com/astronely/financial-helper_microservices/boardService/internal/repository/redis/board"
	"github.com/astronely/financial-helper_microservices/boardService/internal/service"
	boardService "github.com/astronely/financial-helper_microservices/boardService/internal/service/board"
	"github.com/astronely/financial-helper_microservices/boardService/pkg/client/cache"
	cacheClient "github.com/astronely/financial-helper_microservices/boardService/pkg/client/cache/redis"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db/pg"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db/transaction"
	"github.com/redis/go-redis/v9"
)

type serviceProvider struct {
	grpcConfig  config.GRPCConfig
	pgConfig    config.PGConfig
	redisConfig config.RedisConfig
	tokenConfig config.TokenConfig

	rdb         *redis.Client
	redisClient cache.RedisClient
	dbClient    db.Client
	txManager   db.TxManager

	boardService service.BoardService

	boardRepository      repository.BoardRepository
	boardRedisRepository repository.BoardRedisRepository

	boardImpl *board.Implementation
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

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			panic(err)
		}
		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) TokenConfig() config.TokenConfig {
	if s.tokenConfig == nil {
		cfg, err := env.NewTokenConfig()
		if err != nil {
			panic("Cannot load Token Config" + err.Error())
		}
		s.tokenConfig = cfg
	}

	return s.tokenConfig
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

func (s *serviceProvider) Rdb(ctx context.Context) *redis.Client {
	if s.rdb == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:            s.RedisConfig().Address(),
			MaxIdleConns:    s.RedisConfig().MaxIdle(),
			ConnMaxIdleTime: s.RedisConfig().IdleTimeout(),
		})

		s.rdb = rdb
	}

	return s.rdb
}

func (s *serviceProvider) RedisClient(ctx context.Context) cache.RedisClient {
	if s.redisClient == nil {
		redisClient := cacheClient.NewClient(s.Rdb(ctx), s.RedisConfig())
		s.redisClient = redisClient
	}

	return s.redisClient
}

func (s *serviceProvider) BoardService(ctx context.Context) service.BoardService {
	if s.boardService == nil {
		s.boardService = boardService.NewService(s.BoardRepository(ctx), s.BoardRedisRepository(ctx), s.TxManager(ctx), s.TokenConfig())
	}
	return s.boardService
}

func (s *serviceProvider) BoardRepository(ctx context.Context) repository.BoardRepository {
	if s.boardRepository == nil {
		s.boardRepository = boardRepository.NewRepository(s.DBClient(ctx))
	}
	return s.boardRepository
}

func (s *serviceProvider) BoardRedisRepository(ctx context.Context) repository.BoardRedisRepository {
	if s.boardRedisRepository == nil {
		s.boardRedisRepository = boardRedisRepository.NewRepository(s.RedisClient(ctx))
	}
	return s.boardRedisRepository
}

func (s *serviceProvider) BoardImpl(ctx context.Context) *board.Implementation {
	if s.boardImpl == nil {
		s.boardImpl = board.NewImplementation(s.BoardService(ctx))
	}
	return s.boardImpl
}
