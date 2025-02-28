package config

import "github.com/rezaAmiri123/ormus/adapter/redis"

type Config struct {
	Redis redis.Config `koanf:"redis"`
}
