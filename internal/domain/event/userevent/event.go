package userevent

type Event string
type Queue string

var QueuePattern Queue = "user.*"

const (
	UserCreated Event = "user.event.created"
	UserUpdated Event = "user.event.updated"
	UserDeleted Event = "user.event.deleted"
)

var UserEvents = []Event{
	UserCreated, UserUpdated, UserDeleted,
}
