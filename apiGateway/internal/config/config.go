package config

import "github.com/joho/godotenv"

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type GRPCConfig interface {
	AuthAddress() string
	UserAddress() string
	FinanceAddress() string
	NoteAddress() string
	BoardAddress() string
}

type HTTPConfig interface {
	Address() string
}

type SwaggerConfig interface {
	Address() string
}
