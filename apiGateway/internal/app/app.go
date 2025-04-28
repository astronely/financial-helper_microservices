package app

import (
	"context"
	"errors"
	"flag"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config"
	descAccess "github.com/astronely/financial-helper_microservices/apiGateway/pkg/access_v1"
	descAuth "github.com/astronely/financial-helper_microservices/apiGateway/pkg/auth_v1"
	descBoard "github.com/astronely/financial-helper_microservices/apiGateway/pkg/board_v1"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/closer"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	descNote "github.com/astronely/financial-helper_microservices/apiGateway/pkg/note_v1"
	descTransaction "github.com/astronely/financial-helper_microservices/apiGateway/pkg/transaction_v1"
	descUser "github.com/astronely/financial-helper_microservices/apiGateway/pkg/user_v1"
	descWallet "github.com/astronely/financial-helper_microservices/apiGateway/pkg/wallet_v1"
	_ "github.com/astronely/financial-helper_microservices/apiGateway/statik"

	//descWallet "github.com/astronely/financial-helper_microservices/financeService/pkg/transaction_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
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
	httpServer      *http.Server
	swaggerServer   *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
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
		err := a.runHttpServer(ctx)
		if err != nil {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	go func() {
		err := a.runSwaggerServer(ctx)
		if err != nil {
			log.Fatalf("failed to start swagger: %v", err)
		}
	}()

	<-ctx.Done()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
		a.initSwaggerServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
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

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := descUser.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().UserAddress(), opts)
	if err != nil {
		return err
	}
	err = descAuth.RegisterAuthV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().AuthAddress(), opts)
	if err != nil {
		return err
	}
	err = descAccess.RegisterAccessV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().AuthAddress(), opts)
	if err != nil {
		return err
	}
	err = descWallet.RegisterWalletV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().FinanceAddress(), opts)
	if err != nil {
		return err
	}
	err = descTransaction.RegisterTransactionV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().FinanceAddress(), opts)
	if err != nil {
		return err
	}
	err = descNote.RegisterNoteV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().NoteAddress(), opts)
	if err != nil {
		return err
	}
	err = descBoard.RegisterBoardV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcConfig().BoardAddress(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9090"}, // TODO: Дописать IP клиента
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "Content-Length"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HttpConfig().Address(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) runHttpServer(ctx context.Context) error {
	logger.Info("HTTP server is running",
		"address", a.serviceProvider.HttpConfig().Address(),
	)

	closer.Add(func() error {
		if err := a.httpServer.Shutdown(context.Background()); err == nil {
			logger.Info("HTTP server shutdown gracefully")
		}
		return nil
	})

	err := a.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer(_ context.Context) error {
	logger.Info("Swagger server is running",
		"address:", a.serviceProvider.SwaggerConfig().Address(),
	)

	closer.Add(func() error {
		if err := a.swaggerServer.Shutdown(context.Background()); err == nil {
			logger.Info("Swagger server shutdown gracefully")
		}
		return nil
	})

	err := a.swaggerServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/userApi.swagger.json", serveSwaggerFile("/userApi.swagger.json"))
	mux.HandleFunc("/authApi.swagger.json", serveSwaggerFile("/authApi.swagger.json"))
	mux.HandleFunc("/financeApi.swagger.json", serveSwaggerFile("/financeApi.swagger.json"))
	mux.HandleFunc("/boardApi.swagger.json", serveSwaggerFile("/boardApi.swagger.json"))
	mux.HandleFunc("/noteApi.swagger.json", serveSwaggerFile("/noteApi.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceProvider.SwaggerConfig().Address(),
		Handler: mux,
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//logger.Debug("Serving swagger file",
		//	"file", path)

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//log.Printf("Open swagger file: %s", path)

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		//log.Printf("Read swagger file: %s", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//log.Printf("Write swagger file: %s", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//log.Printf("Served swagger file: %s", path)
	}
}
