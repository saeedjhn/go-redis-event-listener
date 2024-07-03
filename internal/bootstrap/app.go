package bootstrap

import (
	"github.com/saeedjhn/go-redis-event-listener/configs"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/logger"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/db/mysql"
)

type Application struct {
	Config      *configs.Config
	Logger      *logger.Logger
	MysqlDB     mysql.DB
	RedisClient redis.DB
}

func App(env configs.Env) *Application {
	var app = &Application{}
	app.Config = ConfigLoad(env)
	app.Logger = NewLogger(app.Config.Logger)
	app.MysqlDB = NewMysqlConnection(app.Config.Mysql)
	app.RedisClient = NewRedisClient(app.Config.Redis)

	return app
}

//func (a *Application) ClosePostgresqlConnection() {
//	ClosePostgresConnection(a.PostgresDB)
//}

func (a *Application) CloseMysqlConnection() {
	CloseMysqlConnection(a.MysqlDB)
}

func (a *Application) CloseRedisClientConnection() {
	CloseRedisClient(a.RedisClient)
}
