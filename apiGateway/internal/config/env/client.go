package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config"
	"net"
	"os"
)

const (
	clientHostEnvName       = "CLIENT_HOST"
	clientPortEnvName       = "CLIENT_PORT"
	clientConnectionEnvName = "CLIENT_CONNECTION"
)

var _ config.ClientConfig = (*clientConfig)(nil)

type clientConfig struct {
	host       string
	port       string
	connection string
}

func NewClientConfig() (config.ClientConfig, error) {
	host := os.Getenv(clientHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found in env")
	}
	port := os.Getenv(clientPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found in env")
	}
	connection := os.Getenv(clientConnectionEnvName)
	if len(connection) == 0 {
		return nil, errors.New("http connection not found in env")
	}

	return &clientConfig{
		host:       host,
		port:       port,
		connection: connection,
	}, nil
}

func (c *clientConfig) Address() string {
	return c.connection + "://" + net.JoinHostPort(c.host, c.port)
}
