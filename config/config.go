package config

import (
	"github.com/rezaAmiri123/ormus/adapter/redis"
	"github.com/rezaAmiri123/ormus/manager"
	"github.com/rezaAmiri123/ormus/source"
)

type Config struct {
	Manager manager.Config `koanf:"manager"`
	Redis   redis.Config   `koanf:"redis"`
	Source  source.Config  `koanf:"source"`
}
