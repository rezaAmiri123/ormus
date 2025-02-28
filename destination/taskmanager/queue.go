package taskmanager

import "github.com/rezaAmiri123/ormus/destination/entity/taskentity"

// Queue is an interface for a queue of tasks in the task manager.
type Queue interface {
	Enqueue(task taskentity.Task) error
}
