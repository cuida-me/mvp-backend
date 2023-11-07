package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type client struct {
	client *redis.Client
}

func New() *client {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Endereço do servidor Redis.
		Password: "",               // Senha, se aplicável.
		DB:       0,                // Banco de dados, 0 por padrão.
	})
	return &client{
		client: c,
	}
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := c.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
