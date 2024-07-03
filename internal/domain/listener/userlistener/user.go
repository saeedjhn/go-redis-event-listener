package userlistener

import (
	"context"
	"fmt"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/event/userevent"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
	"log"
)

type UserListener struct {
	ps *pubsub.PubSub
}

func New(ps *pubsub.PubSub) *UserListener {
	return &UserListener{ps: ps}
}

func (l UserListener) Handler(ctx context.Context, queuePattern userevent.Queue) error {
	subscription := *l.ps.PSubscribe(ctx, string(queuePattern))
	defer subscription.Close()

	channel := subscription.Channel()
	// Or just usage from below line
	//channel := l.ps.PSubscribeMessages(ctx, string(queuePattern))

	for msg := range channel {
		switch userevent.Event(msg.Channel) {
		case userevent.UserCreated:
			l.createdHandler(string(userevent.UserCreated), msg.Payload)
		case userevent.UserUpdated:
			l.updatedHandler(string(userevent.UserCreated), msg.Payload)
		case userevent.UserDeleted:
			l.deletedHandler(string(userevent.UserCreated), msg.Payload)
		default:
			return fmt.Errorf("invalid queue: %s", msg.Channel)
		}
	}

	return nil
}

func (l UserListener) createdHandler(queueName string, message string) {
	var u entity.User
	err := l.ps.Unmarshal(message, &u)
	if err != nil {
		fmt.Errorf("[%s] %v", queueName, err)
	}

	// Continue with your logic here:
	log.Println("user created handler running", u)
}

func (l UserListener) updatedHandler(queueName string, message string) {
	// Continue with your logic here:

	log.Println("user updated handler running")
}

func (l UserListener) deletedHandler(queueName string, message string) {
	// Continue with your logic here:

	log.Println("user deleted handler running")
}
