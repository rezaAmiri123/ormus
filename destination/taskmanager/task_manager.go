package taskmanager

import "github.com/rezaAmiri123/ormus/destination/entity"

type TaskManager interface {
	publish(task entity.Task) error
}
