package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/userService/internal/config"
	"github.com/astronely/financial-helper_microservices/userService/pkg/logger"
	"os"
	"time"
)

const (
	accessTokenKeyEnv             = "ACCESS_TOKEN_KEY"
	refreshTokenKeyEnv            = "REFRESH_TOKEN_KEY"
	accessTokenExpirationTimeEnv  = "ACCESS_TOKEN_EXPIRATION_TIME"
	refreshTokenExpirationTimeEnv = "REFRESH_TOKEN_EXPIRATION_TIME"
)

var _ config.TokenConfig = (*tokenConfig)(nil)

type tokenConfig struct {
	accessTokenKey             string
	refreshTokenKey            string
	accessTokenExpirationTime  time.Duration
	refreshTokenExpirationTime time.Duration
}

func NewTokenConfig() (config.TokenConfig, error) {
	accessTokenKey := os.Getenv(accessTokenKeyEnv)
	if len(accessTokenKey) == 0 {
		return nil, errors.New("access token key not found in environment")
	}

	refreshTokenKey := os.Getenv(refreshTokenKeyEnv)
	if len(refreshTokenKey) == 0 {
		return nil, errors.New("refresh token key not found in environment")
	}

	accessTokenExpirationTime := os.Getenv(accessTokenExpirationTimeEnv)
	if len(accessTokenExpirationTime) == 0 {
		return nil, errors.New("access token expiration time not found in environment")
	}

	accessTokenExpirationTimeDuration, err := time.ParseDuration(accessTokenExpirationTime)
	if err != nil {
		logger.Error("access token expiration time incorrect in environment",
			"error", err)
		accessTokenExpirationTimeDuration = time.Minute * 5
	}

	refreshTokenExpirationTime := os.Getenv(refreshTokenExpirationTimeEnv)
	if len(refreshTokenExpirationTime) == 0 {
		return nil, errors.New("refresh token expiration time not found in environment")
	}

	refreshTokenExpirationTimeDuration, err := time.ParseDuration(refreshTokenExpirationTime)
	if err != nil {
		logger.Error("refresh token expiration time incorrect in environment",
			"error", err)
		refreshTokenExpirationTimeDuration = time.Minute * 15
	}

	return &tokenConfig{
		accessTokenKey:             accessTokenKey,
		refreshTokenKey:            refreshTokenKey,
		accessTokenExpirationTime:  accessTokenExpirationTimeDuration,
		refreshTokenExpirationTime: refreshTokenExpirationTimeDuration,
	}, nil
}

func (cfg *tokenConfig) AccessTokenKey() string {
	return cfg.accessTokenKey
}

func (cfg *tokenConfig) RefreshTokenKey() string {
	return cfg.refreshTokenKey
}

func (cfg *tokenConfig) AccessTokenExpirationTime() time.Duration {
	return cfg.accessTokenExpirationTime
}

func (cfg *tokenConfig) RefreshTokenExpirationTime() time.Duration {
	return cfg.refreshTokenExpirationTime
}
