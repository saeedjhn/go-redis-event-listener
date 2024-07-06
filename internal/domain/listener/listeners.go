package listener

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
)

type Listener struct {
	ps pubsub.PubSub
}

func New(ps pubsub.PubSub) *Listener {
	return &Listener{ps: ps}
}

func (l Listener) SetupListeners() {
	// Todo - Add all listener
}
