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

type HTTPConfig interface {
	Address() string
}

type SwaggerConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type TokenConfig interface {
	AccessTokenKey() string
	RefreshTokenKey() string
	AccessTokenExpirationTime() time.Duration
	RefreshTokenExpirationTime() time.Duration
}
