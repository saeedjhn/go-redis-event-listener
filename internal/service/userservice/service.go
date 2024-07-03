package userservice

import (
	"context"
	"github.com/saeedjhn/go-redis-event-listener/configs"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
)

type AuthGenerator interface {
	CreateAccessToken(dto userauthservicedto.CreateTokenRequest) (userauthservicedto.CreateTokenResponse, error)
	CreateRefreshToken(dto userauthservicedto.CreateTokenRequest) (userauthservicedto.CreateTokenResponse, error)
	ExtractIDFromRefreshToken(dto userauthservicedto.ExtractIDFromTokenRequest) (userauthservicedto.ExtractIDFromTokenResponse, error)
}

type TaskGenerator interface {
	Create(dto usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(dto usertaskservicedto.TasksUserRequest) (usertaskservicedto.TasksUserResponse, error)
}

type Event interface {
	Created(ctx context.Context, u entity.User) error
	Updated(ctx context.Context, id uint) error
	Deleted(ctx context.Context, u entity.User) error
}

type Repository interface {
	Create(u entity.User) (entity.User, error)
	IsMobileUnique(mobile string) (bool, error)
	GetByMobile(mobile string) (entity.User, error)
	GetByID(id uint) (entity.User, error)
}

type UserInteractor struct {
	config         *configs.Config
	redisClient    redis.DB
	authInteractor AuthGenerator
	taskInteractor TaskGenerator
	event          Event
	repository     Repository
}

func New(
	config *configs.Config,
	redisClient redis.DB,
	authInteractor AuthGenerator,
	taskInteractor TaskGenerator,
	event Event,
	repository Repository,
) *UserInteractor {
	return &UserInteractor{
		config:         config,
		redisClient:    redisClient,
		authInteractor: authInteractor,
		taskInteractor: taskInteractor,
		event:          event,
		repository:     repository,
	}
}
