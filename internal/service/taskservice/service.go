package taskservice

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
)

type Repository interface {
	Create(u entity.Task) (entity.Task, error)
	GetByID(id uint) (entity.Task, error)
	GetAllByUserID(userID uint) ([]entity.Task, error)
	IsExistsUser(id uint) (bool, error)
	// etc
}

type TaskInteractor struct {
	redisClient redis.DB
	repository  Repository
}

func New(redisClient redis.DB, repository Repository) *TaskInteractor {
	return &TaskInteractor{redisClient: redisClient, repository: repository}
}
