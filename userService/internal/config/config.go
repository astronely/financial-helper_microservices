package config

import (
	"github.com/joho/godotenv"
	"time"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type GRPCConfig interface {
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
