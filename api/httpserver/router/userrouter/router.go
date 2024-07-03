package userrouter

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-redis-event-listener/api/httpserver/handler/userhandler"
	"github.com/saeedjhn/go-redis-event-listener/api/httpserver/middleware"
	"github.com/saeedjhn/go-redis-event-listener/internal/bootstrap"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/event/userevent"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/token"
	"github.com/saeedjhn/go-redis-event-listener/internal/repository/taskrepository/mysqltask"
	"github.com/saeedjhn/go-redis-event-listener/internal/repository/userrespository/mysqluser"
	"github.com/saeedjhn/go-redis-event-listener/internal/service/authservice"
	"github.com/saeedjhn/go-redis-event-listener/internal/service/taskservice"
	"github.com/saeedjhn/go-redis-event-listener/internal/service/userservice"
	"github.com/saeedjhn/go-redis-event-listener/internal/validator/uservalidator"
)

func New(
	app *bootstrap.Application,
	group *echo.Group,
) {
	// Repository
	taskMysql := mysqltask.New(app.MysqlDB)
	userMysql := mysqluser.New(app.MysqlDB)

	// Event
	//publisher := pubsub.NewPublisher(app.RedisClient)
	ps := pubsub.New(app.RedisClient)
	userEvent := userevent.New(ps)

	// Usecase
	taskCase := taskservice.New(app.RedisClient, taskMysql)
	authCase := authservice.New(app.Config.Auth, token.New())

	// Service-oriented - inject service to another service
	// Repository & Usecase
	userCase := userservice.New(
		app.Config, app.RedisClient, authCase, taskCase, userEvent, userMysql,
	)

	// Validator
	validator := uservalidator.New()

	// Handler
	handler := userhandler.New(app, validator, userCase)

	usersGroup := group.Group("/users")
	{
		publicRouter := usersGroup.Group("")
		{
			publicRouter.POST("/refresh-token", handler.RefreshToken)
		}

		authRouter := usersGroup.Group("/auth")
		{
			authRouter.POST("/register", handler.Register)
			authRouter.POST("/login", handler.Login)
		}

		protectedRouter := usersGroup.Group("")
		//protectedRouter.Use(middleware.Auth(app.Config.Auth, authCase))
		{
			protectedRouter.GET("/profile", handler.Profile)
			protectedRouter.POST("/:id/tasks", handler.CreateTask, middleware.CheckIsValidUserID)
			protectedRouter.GET("/:id/tasks", handler.Tasks)
		}
	}
}
