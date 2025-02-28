package integrationhandler

import (
	"github.com/rezaAmiri123/ormus/destination/entity/taskentity"
	"github.com/rezaAmiri123/ormus/destination/integrationhandler/param"
	"github.com/rezaAmiri123/ormus/event"
)

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(task taskentity.Task, processedEvent event.ProcessedEvent) (param.HandleTaskResponse, error)
}
