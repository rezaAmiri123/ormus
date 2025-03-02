package taskcoordinator

import (
	"sync"

	"github.com/rezaAmiri123/ormus/event"
)

type Coordinator interface {
	Start(processedEvents <-chan event.ProcessedEvent, done <-chan bool, wg *sync.WaitGroup)
}
