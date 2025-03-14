package tests

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/rezaAmiri123/ormus/adapter/redis"
	scylladbsession "github.com/rezaAmiri123/ormus/adapter/scylladb"
	"github.com/rezaAmiri123/ormus/config"
	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/pkg/channel"
	"github.com/rezaAmiri123/ormus/pkg/channel/adapter/rabbitmqchannel"
	"github.com/rezaAmiri123/ormus/source/adapter/manager"
	sourceevent "github.com/rezaAmiri123/ormus/source/eventhandler"
	writekeyrepo "github.com/rezaAmiri123/ormus/source/repository/redis/rediswritekey"
	"github.com/rezaAmiri123/ormus/source/repository/scylladb"
	eventrepo "github.com/rezaAmiri123/ormus/source/repository/scylladb/event"
	eventsvc "github.com/rezaAmiri123/ormus/source/service/event"
	"github.com/rezaAmiri123/ormus/source/service/writekey"
)

type TestConfig struct {
	EventRepo     eventrepo.Repository
	Cfg           config.Config
	SylladbConn   scylladbsession.SessionxInterface
	brokerAdapter *rabbitmqchannel.ChannelAdapter
	Consumer      sourceevent.Consumer
	writeKeyRepo  writekeyrepo.DB
	redisAdapter  redis.Adapter
}

func ClearUp() {
	c := GetConfigs()
	err := c.SylladbConn.ExecStmt("TRUNCATE test_source.event")
	if err != nil {
		logger.L().Error(err.Error())
	}
	_, err = c.brokerAdapter.PurgeTheChannel(c.Cfg.Source.NewEventQueueName)
	if err != nil {
		logger.L().Error(err.Error())
	}
}

func TestMain(m *testing.M) {
	ClearUp()
	exitVal := m.Run()
	ClearUp()
	os.Exit(exitVal)
}

func GetConfigs() TestConfig {
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	cfg := config.New(config.Option{
		Prefix:       "ORMUS_",
		Delimiter:    ".",
		Separator:    "__",
		YamlFilePath: fmt.Sprintf("%s/test_config.yml", os.Getenv("ROOT")),
	})
	Adapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err := Adapter.NewChannel(cfg.Source.NewSourceEventName, channel.BothMode, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}
	err = Adapter.NewChannel(cfg.Source.NewEventQueueName, channel.BothMode, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	Publisher := sourceevent.NewPublisher(Adapter)
	DB, err := scylladb.New(cfg.Source.ScyllaDBConfig)
	if err != nil {
		panic(err)
	}

	repo := eventrepo.New(DB, Publisher)
	redisAdapter, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}
	ManagerAdapter := manager.New(cfg.Source)
	writeKeyRepo := writekeyrepo.New(redisAdapter, *ManagerAdapter)
	writeKeySvc := writekey.New(&writeKeyRepo, cfg.Source)
	eventRepo := eventrepo.New(DB, Publisher)
	eventSvc := *eventsvc.New(eventRepo, cfg.Source, wg)
	Consumer := *sourceevent.NewConsumer(Adapter, writeKeySvc, eventSvc, cfg.Source.RetryNumber)
	return TestConfig{
		EventRepo:     *repo,
		Cfg:           cfg,
		SylladbConn:   DB.GetConn(),
		brokerAdapter: Adapter,
		Consumer:      Consumer,
		writeKeyRepo:  writeKeyRepo,
		redisAdapter:  redisAdapter,
	}
}
