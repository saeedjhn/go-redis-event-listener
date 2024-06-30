package pubsub

import (
	"context"
	"encoding/json"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"log"
)

// Publisher is a generic publisher for different message types.
type Publisher struct {
	redisClient redis.DB
}

func NewPublisher(redisClient redis.DB) *Publisher {
	return &Publisher{redisClient}
}

// PublishMessages publishes messages to  channels.
func (p *Publisher) PublishMessages(ctx context.Context, message interface{}, queueName string) {

	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("[%s] Failed to serialize message: %v", queueName, err)
	}

	// Use the context for the publishing operation
	err = p.redisClient.Client().Publish(ctx, queueName, serializedMessage).Err()
	if err != nil {
		log.Printf("[%s] Failed to publish message: %v", queueName, err)
	}

}
