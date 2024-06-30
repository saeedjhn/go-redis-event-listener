package bootstrap

import "github.com/saeedjhn/go-redis-pubsub-message-broker/configs"

func ConfigLoad(env configs.Env) *configs.Config {
	return configs.Load(env)
}
