package userdto

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"time"
)

type TaskInfo struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      entity.Status `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
