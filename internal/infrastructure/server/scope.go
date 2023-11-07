package server

import (
	"github.com/cuida-me/mvp-backend/pkg/env"
	"time"
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
	Environment  Environment
	LogLevel     string
	Port         string
	Network      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Shutdown func() error

type Scope interface {
	Bootstrap(cfg Config) error
}

func NewConfig() *Config {
	return &Config{
		Environment:  Environment(env.GetString("SCOPE", "local")),
		Port:         env.GetString("PORT", ":8080"),
		Network:      env.GetString("NETWORK", "tcp"),
		LogLevel:     env.GetString("LOG_LEVEL", "INFO"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
