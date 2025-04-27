package env

import (
	"fmt"
	"github.com/astronely/financial-helper_microservices/boardService/internal/config"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	redisHostEnv                 = "REDIS_HOST"
	redisPortEnv                 = "REDIS_PORT"
	redisConnectionTimeoutSecEnv = "REDIS_CONNECTION_TIMEOUT_SEC"
	redisMaxIdleEnv              = "REDIS_MAX_IDLE"
	redisIdleTimeoutSecEnv       = "REDIS_IDLE_TIMEOUT_SEC"
)

type redisConfig struct {
	host string
	port string

	connectionTimeout time.Duration

	maxIdle     int
	idleTimeout time.Duration
}

func NewRedisConfig() (config.RedisConfig, error) {
	host := os.Getenv(redisHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s not set", redisHostEnv)
	}

	port := os.Getenv(redisPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s not set", redisPortEnv)
	}

	connectionTimeoutSec, err := strconv.ParseInt(os.Getenv(redisConnectionTimeoutSecEnv), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s not set", redisConnectionTimeoutSecEnv)
	}

	maxIdle, err := strconv.Atoi(os.Getenv(redisMaxIdleEnv))
	if err != nil {
		return nil, fmt.Errorf("environment variable %s not set", redisMaxIdleEnv)
	}

	idleTimeoutSec, err := strconv.ParseInt(os.Getenv(redisIdleTimeoutSecEnv), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s not set", redisIdleTimeoutSecEnv)
	}

	return &redisConfig{
		host:              host,
		port:              port,
		connectionTimeout: time.Duration(connectionTimeoutSec) * time.Second,
		maxIdle:           maxIdle,
		idleTimeout:       time.Duration(idleTimeoutSec) * time.Second,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) ConnectionTimeout() time.Duration {
	return cfg.connectionTimeout
}

func (cfg *redisConfig) MaxIdle() int {
	return cfg.maxIdle
}

func (cfg *redisConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeout
}
