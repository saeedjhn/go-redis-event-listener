package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
)

// Publisher is a generic publisher for different Message types.
type Publisher struct {
	redisClient redis.DB
}

func NewPublisher(redisClient redis.DB) *Publisher {
	return &Publisher{redisClient}
}

// Publish publishes messages to  channels.
func (p *Publisher) Publish(ctx context.Context, message interface{}, queueName string) error {
	serializedMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("[%s] Failed to serialize Message: %v", queueName, err)
	}

	// Use the context for the publishing operation
	err = p.redisClient.Client().Publish(ctx, queueName, serializedMessage).Err()
	if err != nil {
		return fmt.Errorf("[%s] Failed to publish Message: %v", queueName, err)
	}

	return nil
}
