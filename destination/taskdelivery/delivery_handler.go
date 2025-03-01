package taskdelivery

import (
	"github.com/rezaAmiri123/ormus/destination/entity/taskentity"
	"github.com/rezaAmiri123/ormus/destination/taskdelivery/param"
	"github.com/rezaAmiri123/ormus/event"
)

// DeliveryHandler is responsible for delivering event to third party destinations.
type DeliveryHandler interface {
	Handle(task taskentity.Task, processedEvent event.ProcessedEvent) (param.DeliveryTaskResponse, error)
}
