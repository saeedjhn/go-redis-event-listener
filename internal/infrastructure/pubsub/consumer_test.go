package pubsub

import (
	"context"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"testing"
)

func TestConsumer(t *testing.T) {
	r := redis.New(redis.Config{
		Host:               "127.0.0.1",
		Port:               "6379",
		Password:           "123456",
		DB:                 0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		PoolTimeout:        0,
		IdleCheckFrequency: 0,
	})

	s := NewConsumer(r)

	s.ConsumerMessages(context.Background(), []string{"Test"})
}
