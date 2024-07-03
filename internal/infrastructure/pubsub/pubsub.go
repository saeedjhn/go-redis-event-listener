package pubsub

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	redisDB "github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
	"sync"
)

var instance *RedisPubSub
var once sync.Once

type RedisPubSub struct {
	client redisDB.DB
}

func New(client redisDB.DB) *RedisPubSub {
	once.Do(func() {
		instance = &RedisPubSub{client: client}
	})

	return instance
}

func (p *RedisPubSub) Publish(ctx context.Context, message interface{}, queueName string) error {
	serializedMessage, err := p.marshal(message)
	if err != nil {
		return fmt.Errorf("[%s] %v", queueName, err)
	}

	encodeMessage := p.base64EncodeToString(serializedMessage)

	// Use the context for the publishing operation
	err = p.client.Client().Publish(ctx, queueName, encodeMessage).Err()
	if err != nil {
		return fmt.Errorf("[%s] Failed to publish Message: %v", queueName, err)
	}

	return nil
}

func (p *RedisPubSub) Decode(message string, ptr interface{}) error {
	payload, err := p.base64DecodeString(message)
	if err != nil {
		return err
	}

	if err = p.unmarshal(payload, ptr); err != nil {
		return err
	}

	return nil
}

func (p *RedisPubSub) PSubscribe(ctx context.Context, queuePattern string) redisDB.PubSub {
	return p.client.Client().PSubscribe(ctx, queuePattern)
}

func (p *RedisPubSub) PSubscribeMessages(ctx context.Context, queuePattern string) redisDB.Message {
	subscription := p.client.Client().PSubscribe(ctx, queuePattern)

	return subscription.Channel()
}

func (p *RedisPubSub) Subscribe(ctx context.Context, queue string) redisDB.PubSub {
	return p.client.Client().Subscribe(ctx, queue)
}

func (p *RedisPubSub) SubscribeMessages(ctx context.Context, queue string) redisDB.Message {
	subscription := p.client.Client().Subscribe(ctx, queue)

	return subscription.Channel()
}

func (p *RedisPubSub) base64EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func (p *RedisPubSub) base64DecodeString(message string) ([]byte, error) {
	payload, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to decode Message: %v", err)
	}

	return payload, nil
}

func (p *RedisPubSub) marshal(message interface{}) ([]byte, error) {
	mSerialized, err := json.Marshal(message)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to serialize Message: %v", err)
	}

	return mSerialized, nil
}

func (p *RedisPubSub) unmarshal(payload []byte, ptr interface{}) error {
	err := json.Unmarshal(payload, ptr)
	if err != nil {
		return fmt.Errorf("failed to deserialize message: %v", err)
	}
	return nil
}
