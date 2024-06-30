package main

import (
	"context"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/pubsub"
	"log"
	"time"
)

type Data struct {
	Foo string
	Bar string
}

func main() {
	//app := bootstrap.App(configs.Development)

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

	//consumer := pubsub.NewConsumer(app.RedisClient)
	log.Println("consumer is running")
	consumer := pubsub.NewConsumer(r)
	consumer.ConsumerMessages(ctx, []string{queue})

	//time.Sleep(1 * time.Second)
	//publisher := pubsub.NewPublisher(app.RedisClient)
	log.Println("publisher is running")
	publisher := pubsub.NewPublisher(r)
	time.Sleep(2 * time.Second)
	data := Data{
		Foo: "foo",
		Bar: "bar",
	}
	publisher.PublishMessages(ctx, data, queue)

	time.Sleep(50 * time.Second)
}
