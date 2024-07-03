package listener

import "github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"

type Listener struct {
	client redis.DB
}

func New(client redis.DB) *Listener {
	return &Listener{client: client}
}

func (l Listener) SetupListeners() {
	// Todo - Add all listener
}
