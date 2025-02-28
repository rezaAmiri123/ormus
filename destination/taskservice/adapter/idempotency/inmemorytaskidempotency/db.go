package inmemorytaskidempotency

import (
	"sync"

	"github.com/rezaAmiri123/ormus/destination/entity/taskentity"
)

type DB struct {
	tasks map[string]taskentity.IntegrationDeliveryStatus
	mutex *sync.Mutex
}

func New() DB {
	tasks := make(map[string]taskentity.IntegrationDeliveryStatus)

	return DB{
		tasks: tasks,
		mutex: &sync.Mutex{},
	}
}
