package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/boardService/internal/config"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("dsn variable not found in env")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
