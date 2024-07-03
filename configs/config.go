package configs

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/logger"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/cache/redis"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/persistance/db/mysql"
	"github.com/saeedjhn/go-redis-event-listener/internal/service/authservice"
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
	Application Application        `mapstructure:"application"`
	HTTPServer  HTTPServer         `mapstructure:"http_server"`
	Logger      logger.Config      `mapstructure:"logger"`
	Mysql       mysql.Config       `mapstructure:"mysql"`
	Redis       redis.Config       `mapstructure:"redis"`
	Auth        authservice.Config `mapstructure:"auth"`
}
