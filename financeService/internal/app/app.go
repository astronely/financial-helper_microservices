package app

import (
	"context"
	"errors"
	"flag"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/closer"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/financeService/internal/config"
	"github.com/astronely/financial-helper_microservices/financeService/internal/interceptor"
	descTransaction "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	descWallet "github.com/astronely/financial-helper_microservices/financeService/pkg/wallet_v1"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "local.env", "path to config file")
	logger.Init(configPath)
}

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) *App {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		panic(err)
	}

	return a
}

func (a *App) Run() error {
	defer func() {
		logger.Info("shutting down the server...")

		closer.CloseAll()
		closer.Wait()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	<-ctx.Done()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	flag.Parse()
	err := config.Load(configPath)
	return err
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.LogInterceptor,
				interceptor.ValidateInterceptor,
			),
		),
	)

	reflection.Register(a.grpcServer)
	descWallet.RegisterWalletV1Server(a.grpcServer, a.serviceProvider.WalletImpl(ctx))
	descTransaction.RegisterTransactionV1Server(a.grpcServer, a.serviceProvider.TransactionImpl(ctx))
	return nil
}

func (a *App) runGRPCServer() error {
	logger.Info("GRPC server is running",
		"address:", a.serviceProvider.GRPCConfig().Address(),
	)

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	closer.Add(func() error {
		a.grpcServer.GracefulStop()
		logger.Info("GRPC server shutdown gracefully")
		return nil
	})

	err = a.grpcServer.Serve(list)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
