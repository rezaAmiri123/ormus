package main

import (
	"github.com/rezaAmiri123/ormus/config"
	"github.com/rezaAmiri123/ormus/destination/integrationhandler/adapters/fakeintegrationhandler"
	"github.com/rezaAmiri123/ormus/destination/processedevent/adapter/rabbitmqconsumer"
	"github.com/rezaAmiri123/ormus/destination/taskidemotency/adapter/inmemorytaskrepo"
	"github.com/rezaAmiri123/ormus/destination/taskidemotency/service/taskidempotencyservice"
	"github.com/rezaAmiri123/ormus/destination/taskmanager"
	"github.com/rezaAmiri123/ormus/destination/taskmanager/adapter/inmemorytaskmanager"
	"github.com/rezaAmiri123/ormus/destination/taskmanager/adapter/rabbitmqtaskmanager"
	"log"
	"sync"
)

func main() {

	//-----Setup queue and workers-----
	var workers []taskmanager.Worker

	fh := fakeintegrationhandler.New()

	taskIdempotencyRepo := inmemorytaskrepo.New()
	taskIdempotencySrv := taskidempotencyservice.New(taskIdempotencyRepo)

	// Get connection config for rabbitMQ
	rmqTaskManagerConnConfig := config.C().Destination.RabbitMQTaskManagerConnection

	// Create RabbitMQ Queue and Workers for webhooks integration
	rmqTaskManager := rabbitmqtaskmanager.NewTaskManager(rmqTaskManagerConnConfig, "webhook_queue", taskIdempotencySrv)
	workers = append(workers, rabbitmqtaskmanager.NewWorker(rmqTaskManager, fh))

	// Create In Memory Queue and Workers for webhooks integration
	inMemoryTaskManager := inmemorytaskmanager.NewTaskManager(taskIdempotencySrv)
	workers = append(workers, inmemorytaskmanager.NewWorker(inMemoryTaskManager, fh))

	// Start workers
	wg := sync.WaitGroup{}
	for _, wrk := range workers {
		wg.Add(1)
		go func(w taskmanager.Worker) {
			defer wg.Done()
			w.ProcessJobs()
		}(wrk)
	}

	// Get connection config for rabbitMQ consumer
	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	rmqConsumerTopic := config.C().Destination.ConsumerTopic
	rmqConsumer := rabbitmqconsumer.NewConsumer(rmqConsumerTopic, rmqConsumerConnConfig)

	// Consume processed events
	err := rmqConsumer.Consume()
	if err != nil {
		log.Panicf("Error on consuming processed events")
	}
	// Wait for workers to finish
	wg.Wait()
}
