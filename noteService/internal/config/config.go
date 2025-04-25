package config

import "github.com/joho/godotenv"

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
