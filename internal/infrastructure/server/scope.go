package server

import "github.com/cuida-me/mvp-backend/pkg/env"

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
	AwsSecretKey string
	AwsAccessKey string
	AwsBucket    string
	AwsRegion    string
}

type Shutdown func() error

type Scope interface {
	Bootstrap(cfg Config) error
}

func NewConfig() *Config {
	return &Config{
		Environment:  Environment(env.GetString("SCOPE", "local")),
		Port:         env.GetString("PORT", ":50051"),
		Network:      env.GetString("NETWORK", "tcp"),
		LogLevel:     env.GetString("LOG_LEVEL", "INFO"),
		AwsSecretKey: env.GetString("AWS_SECRET_KEY", ""),
		AwsAccessKey: env.GetString("AWS_ACCESS_KEY", ""),
		AwsBucket:    env.GetString("AWS_BUCKET", ""),
		AwsRegion:    env.GetString("AWS_REGION", ""),
	}
}
