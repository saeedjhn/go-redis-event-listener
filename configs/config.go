package configs

import (
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/db/mysql"
	"time"
)

const (
	Development Env = "development"
	Production  Env = "production"
)

type Application struct {
	Env                     Env           `mapstructure:"env"`
	Debug                   bool          `mapstructure:"debug"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type HTTPServer struct {
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Config struct {
	Application Application  `mapstructure:"application"`
	HTTPServer  HTTPServer   `mapstructure:"http_server"`
	Mysql       mysql.Config `mapstructure:"mysql"`
	Redis       redis.Config `mapstructure:"redis"`
}
