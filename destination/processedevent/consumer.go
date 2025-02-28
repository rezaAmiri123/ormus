package processedevent

import "github.com/rezaAmiri123/ormus/event"

type Consumer interface {
	Consume(done <-chan bool) (<-chan event.ProcessedEvent, error)
	Close() error
}
