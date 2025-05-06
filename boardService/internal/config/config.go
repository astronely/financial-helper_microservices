package config

import (
	"github.com/joho/godotenv"
	"time"
)

func Load(path string) error {
	err := godotenv.Load(path)
	return err
}

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type TokenConfig interface {
	AccessTokenKey() string
	RefreshTokenKey() string
	AccessTokenExpirationTime() time.Duration
	RefreshTokenExpirationTime() time.Duration
}
