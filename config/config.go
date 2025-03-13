package config

import (
	"github.com/rezaAmiri123/ormus/adapter/etcd"
	"github.com/rezaAmiri123/ormus/adapter/redis"
	"github.com/rezaAmiri123/ormus/adapter/scylladb"
	"github.com/rezaAmiri123/ormus/destination/dconfig"
	"github.com/rezaAmiri123/ormus/manager"
	"github.com/rezaAmiri123/ormus/pkg/channel/adapter/rabbitmqchannel"
	"github.com/rezaAmiri123/ormus/source"
)

type Config struct {
	Redis       redis.Config           `koanf:"redis"`
	Etcd        etcd.Config            `koanf:"etcd"`
	RabbitMq    rabbitmqchannel.Config `koanf:"rabbitmq"`
	Manager     manager.Config         `koanf:"manager"`
	Source      source.Config          `koanf:"source"`
	Destination dconfig.Config         `koanf:"destination"`
	Scylladb    scylladb.Config        `koanf:"scylladb"`
	Swagger     Swagger                `koanf:"swagger"`
}
