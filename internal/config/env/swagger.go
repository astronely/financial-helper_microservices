package env

import (
	"errors"
	"net"
	"os"
)

const (
	swaggerHostEnv = "SWAGGER_HOST"
	swaggerPortEnv = "SWAGGER_PORT"
)

type SwaggerConfig interface {
	Address() string
}

type swaggerConfig struct {
	host string
	port string
}

func NewSwaggerConfig() (SwaggerConfig, error) {
	host := os.Getenv(swaggerHostEnv)
	if len(host) == 0 {
		return nil, errors.New("http host not found in environment")
	}

	port := os.Getenv(swaggerPortEnv)
	if len(port) == 0 {
		return nil, errors.New("http port not found in environment")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *swaggerConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
