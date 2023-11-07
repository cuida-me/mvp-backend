package cache

import (
	"context"
	"time"
)

//go:generate mockgen -destination=./mocks.go -package=cache -source=./contracts.go

type Provider interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}
