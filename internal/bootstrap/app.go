package bootstrap

import (
	"github.com/saeedjhn/go-redis-pubsub-message-broker/configs"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/db/mysql"
)

type Application struct {
	Config      *configs.Config
	MysqlDB     mysql.DB
	RedisClient redis.DB
}

func App(env configs.Env) *Application {
	var app = &Application{}
	app.Config = ConfigLoad(env)
	app.MysqlDB = NewMysqlConnection(app.Config.Mysql)
	app.RedisClient = NewRedisClient(app.Config.Redis)

	return app
}

func (a *Application) CloseMysqlConnection() {
	CloseMysqlConnection(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	CloseRedisClient(a.RedisClient)
}
