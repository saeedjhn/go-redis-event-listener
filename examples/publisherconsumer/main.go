package main

import (
	"context"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
	"log"
	"time"
)

func main() {
	r := redis.New(redis.Config{
		Host:               "localhost",
		Port:               "6379",
		Password:           "123456",
		DB:                 0,
		PoolSize:           0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolTimeout:        0,
		IdleCheckFrequency: 0,
	})

	ctx := context.Background()
	queue := "Test"

	consumer := pubsub.NewConsumer(r)
	log.Println("consumer is running")

	consumer.Consumer(ctx, []string{queue})

	publisher := pubsub.NewPublisher(r)
	log.Println("publisher is running")

	// wait for consumer running
	time.Sleep(2 * time.Second)

	publisher.Publish(ctx, "- MSG -", queue)

	time.Sleep(10 * time.Second)
}
