package taskdelivery

import (
	"github.com/rezaAmiri123/ormus/destination/entity/taskentity"
	"github.com/rezaAmiri123/ormus/destination/taskdelivery/param"
)

// DeliveryHandler is responsible for delivering processed event to third party destinations.
type DeliveryHandler interface {
	Handle(task taskentity.Task) (param.DeliveryTaskResponse, error)
}
