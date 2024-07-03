package userevent

import (
	"context"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
)

type UserEvent struct {
	//publisher *pubsub.Publisher
	pubSub *pubsub.RedisPubSub
}

// func New(publisher *pubsub.Publisher) *UserEvent {
func New(pubSub *pubsub.RedisPubSub) *UserEvent {
	return &UserEvent{pubSub: pubSub}
}

func (e UserEvent) Created(ctx context.Context, u entity.User) error {
	return e.publish(ctx, UserCreated, u)
}

func (e UserEvent) Updated(ctx context.Context, id uint) error {
	return e.publish(ctx, UserUpdated, "users.event.updated")
}

func (e UserEvent) Deleted(ctx context.Context, u entity.User) error {
	return e.publish(ctx, UserDeleted, "users.event.deleted")
}

func (e UserEvent) publish(ctx context.Context, queueName Event, message interface{}) error {
	return e.pubSub.Publish(ctx, message, string(queueName))
}
