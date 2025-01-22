package main

import (
	"context"
	"flag"
	_ "github.com/Masterminds/squirrel"
	userAPI "github.com/astronely/financial-helper_microservices/internal/api/user"
	"github.com/astronely/financial-helper_microservices/internal/config"
	"github.com/astronely/financial-helper_microservices/internal/config/env"
	userRepository "github.com/astronely/financial-helper_microservices/internal/repository/user"
	userService "github.com/astronely/financial-helper_microservices/internal/service/user"
	desc "github.com/astronely/financial-helper_microservices/pkg/user_v1"
	_ "github.com/brianvoe/gofakeit/v7"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "local.env", "path to config file")
}

func main() {
	ctx := context.Background()

	flag.Parse()
	err := config.Load(configPath)

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("Failed to load grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("Failed to load pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, userAPI.NewImplementation(userSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
