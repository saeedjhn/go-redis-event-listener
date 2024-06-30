package pubsub

import (
	"context"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"log"
	"testing"
)

func TestPublish(t *testing.T) {
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

	pub := NewPublisher(r)

	q := "test"
	m := "message for test"

	pub.PublishMessages(context.Background(), m, q)

	log.Println("publish message on channel successful")
}
