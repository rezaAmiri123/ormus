package taskmanager

import "github.com/rezaAmiri123/ormus/destination/entity"

// Queue is an interface for a queue of tasks in the task manager.
type Queue interface {
	enqueue(task entity.Task) error
}
