package bootstrap

import "github.com/saeedjhn/go-redis-event-listener/configs"

func ConfigLoad(env configs.Env) *configs.Config {
	return configs.Load(env)
}
