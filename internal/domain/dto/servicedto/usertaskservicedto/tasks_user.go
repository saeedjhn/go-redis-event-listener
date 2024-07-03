package usertaskservicedto

import "github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"

type TasksUserRequest struct {
	UserID uint
}

type TasksUserResponse struct {
	Tasks []entity.Task
}
