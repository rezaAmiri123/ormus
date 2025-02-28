package rabbitmqtaskmanager

import (
	"github.com/rezaAmiri123/ormus/destination/dconfig"
	"github.com/rezaAmiri123/ormus/event"
)

type TaskManager struct {
	queue *Queue
}

func NewTaskManager(c dconfig.RabbitMQTaskManagerConnection, queueName string) *TaskManager {
	q := newQueue(c, queueName)

	return &TaskManager{
		queue: q,
	}
}

func (tm *TaskManager) Publish(pe event.ProcessedEvent) error {
	err := tm.queue.Enqueue(pe)
	if err != nil {
		return err
	}

	return nil
}
