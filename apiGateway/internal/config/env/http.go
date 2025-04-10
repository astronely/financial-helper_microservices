package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHttpConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found in env")
	}
	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found in env")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
