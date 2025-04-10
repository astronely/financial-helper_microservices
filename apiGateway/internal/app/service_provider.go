package app

import (
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config/env"
)

type serviceProvider struct {
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			panic("Error loading grpc config")
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) HttpConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHttpConfig()
		if err != nil {
			panic("Error loading http config")
		}
		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			panic("Error loading swagger config")
		}
		s.swaggerConfig = cfg
	}
	return s.swaggerConfig
}
