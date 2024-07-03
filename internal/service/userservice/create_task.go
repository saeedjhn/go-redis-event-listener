package userservice

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/pkg/message"
)

type Data struct {
	Foo string
	Bar string
}

func (u *UserInteractor) CreateTask(req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error) {
	const op = message.OpUserUsecaseCreateTask

	dto := usertaskservicedto.CreateTaskRequest{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.Pending,
	}

	createdTask, err := u.taskInteractor.Create(dto)
	if err != nil {
		return userdto.CreateTaskResponse{}, err
	}

	return userdto.CreateTaskResponse{Task: createdTask.Task}, nil
}
