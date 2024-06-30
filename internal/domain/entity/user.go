package entity

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Event string

const (
	MatchingUsersMatchedEvent Event = "matching.users_matched"
)
