package pubsub

import (
	"context"
	redisDB "github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
)

type PubSub interface {
	Publish(ctx context.Context, message interface{}, queueName string) error
	Decode(message string, ptr interface{}) error
	PSubscribe(ctx context.Context, queuePattern string) redisDB.PubSub
	PSubscribeMessages(ctx context.Context, queuePattern string) redisDB.Message
	Subscribe(ctx context.Context, queue string) redisDB.PubSub
	SubscribeMessages(ctx context.Context, queue string) redisDB.Message
}
