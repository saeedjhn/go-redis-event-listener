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

	ps := pubsub.New(r)

	log.Println("subscribe is running")
	go func() {
		//pSubscribe := *ps.PSubscribe(ctx, queue)
		//defer pSubscribe.Close()
		//
		//ch := pSubscribe.Channel()
		//for msg := range ch {
		//	log.Println(msg.Payload)
		//}

		for msg := range ps.PSubscribeMessages(ctx, queue) {
			log.Println(msg.Payload)
		}
	}()

	time.Sleep(1 * time.Second)

	log.Println("publish is running")
	ps.Publish(ctx, "- MSG -", queue)
	ps.Publish(ctx, "- MSG -", queue)

	time.Sleep(10 * time.Second)
}
