package app

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/closer"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/noteService/internal/api/note"
	"github.com/astronely/financial-helper_microservices/noteService/internal/config"
	"github.com/astronely/financial-helper_microservices/noteService/internal/config/env"
	"github.com/astronely/financial-helper_microservices/noteService/internal/repository"
	noteRepository "github.com/astronely/financial-helper_microservices/noteService/internal/repository/note"
	"github.com/astronely/financial-helper_microservices/noteService/internal/service"
	noteService "github.com/astronely/financial-helper_microservices/noteService/internal/service/note"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db/pg"
)

type serviceProvider struct {
	grpcConfig  config.GRPCConfig
	pgConfig    config.PGConfig
	tokenConfig config.TokenConfig

	dbClient db.Client

	noteService service.NoteService

	noteRepository repository.NoteRepository

	noteImpl *note.Implementation
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

func (s *serviceProvider) NoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.NoteRepository(ctx), s.TokenConfig())
	}
	return s.noteService
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.DBClient(ctx))
	}
	return s.noteRepository
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteRepository(ctx))
	}
	return s.noteImpl
}
