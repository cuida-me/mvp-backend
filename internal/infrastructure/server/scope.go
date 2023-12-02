package server

import (
	"time"

	"github.com/cuida-me/mvp-backend/pkg/env"
)

type Environment string

func (e Environment) IsProduction() bool {
	return e == PROD
}

func (e Environment) IsBeta() bool {
	return e == BETA
}

func (e Environment) IsLocal() bool {
	return e == LOCAL
}

func (e Environment) IsTest() bool {
	return e == TEST
}

const (
	PROD  Environment = "production"
	BETA  Environment = "beta"
	LOCAL Environment = "local"
	TEST  Environment = "test"
)

type Config struct {
	Environment      Environment
	LogLevel         string
	Port             string
	Network          string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabaseSchema   string
	DatabaseUrl      string
}

type Shutdown func() error

type Scope interface {
	Bootstrap(cfg Config) error
}

func NewConfig() *Config {
	return &Config{
		Environment:      Environment(env.GetString("SCOPE", "local")),
		Port:             env.GetString("PORT", ":8080"),
		Network:          env.GetString("NETWORK", "tcp"),
		LogLevel:         env.GetString("LOG_LEVEL", "INFO"),
		WriteTimeout:     time.Second * 15,
		ReadTimeout:      time.Second * 15,
		IdleTimeout:      time.Second * 60,
		DatabaseUsername: env.GetString("DATABASE_USERNAME", "root"),
		DatabasePassword: env.GetString("DATABASE_PASSWORD", ""),
		DatabaseHost:     env.GetString("DATABASE_HOST", "localhost:3306"),
		DatabaseSchema:   env.GetString("DATABASE_SCHEMA", "cuidamelocal"),
		DatabaseUrl:      env.GetString("DATABASE_URL", ""),
	}
}
