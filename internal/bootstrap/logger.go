package bootstrap

import "github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/logger"

func NewLogger(config logger.Config) *logger.Logger {
	return logger.New(config)
}
