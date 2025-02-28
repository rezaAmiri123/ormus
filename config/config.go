package config

import (
	"github.com/rezaAmiri123/ormus/adapter/redis"
	"github.com/rezaAmiri123/ormus/source"
)

type Config struct {
	Source source.Config `koanf:"source"`
	Redis  redis.Config  `koanf:"redis"`
}