package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	redisDB "github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
	"sync"
)

var instance *PubSub
var once sync.Once

type PubSub struct {
	client redisDB.DB
}

func New(client redisDB.DB) *PubSub {
	once.Do(func() {
		instance = &PubSub{client: client}
	})

	return instance
}

func (p *PubSub) Publish(ctx context.Context, message interface{}, queueName string) error {
	serializedMessage, err := p.marshal(message)
	if err != nil {
		return fmt.Errorf("[%s] %v", queueName, err)
	}

	// Use the context for the publishing operation
	err = p.client.Client().Publish(ctx, queueName, serializedMessage).Err()
	if err != nil {
		return fmt.Errorf("[%s] Failed to publish Message: %v", queueName, err)
	}

	return nil
}

func (p *PubSub) Unmarshal(message string, ptr interface{}) error {
	err := json.Unmarshal([]byte(message), ptr)
	if err != nil {
		return fmt.Errorf("failed to deserialize message: %v", err)
	}

	return nil
}

func (p *PubSub) marshal(message interface{}) ([]byte, error) {
	mSerialized, err := json.Marshal(message)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to serialize Message: %v", err)
	}

	return mSerialized, nil
}

func (p *PubSub) PSubscribe(ctx context.Context, queuePattern string) redisDB.PubSub {
	return p.client.Client().PSubscribe(ctx, queuePattern)
}

func (p *PubSub) PSubscribeMessages(ctx context.Context, queuePattern string) redisDB.Message {
	subscription := p.client.Client().PSubscribe(ctx, queuePattern)

	return subscription.Channel()
}

func (p *PubSub) Subscribe(ctx context.Context, queue string) redisDB.PubSub {
	return p.client.Client().Subscribe(ctx, queue)
}

func (p *PubSub) SubscribeMessages(ctx context.Context, queue string) redisDB.Message {
	subscription := p.client.Client().Subscribe(ctx, queue)

	return subscription.Channel()
}
