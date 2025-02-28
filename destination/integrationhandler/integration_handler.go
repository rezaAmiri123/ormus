package integrationhandler

import (
	"github.com/rezaAmiri123/ormus/event"
)

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(processedEvent event.ProcessedEvent) error
}
