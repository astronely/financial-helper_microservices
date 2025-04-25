package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/noteService/internal/config"
	"net"
	"os"
)

const (
	hostEnvName = "GRPC_HOST"
	portEnvName = "GRPC_PORT"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (config.GRPCConfig, error) {
	host := os.Getenv(hostEnvName)
	if len(host) == 0 {
		return nil, errors.New("host variable not found in env")
	}
	port := os.Getenv(portEnvName)
	if len(port) == 0 {
		return nil, errors.New("port variable not found in env")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
