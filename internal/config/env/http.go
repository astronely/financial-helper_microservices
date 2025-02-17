package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/internal/config"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type HttpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found in environment")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found in environment")
	}

	return &HttpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *HttpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
