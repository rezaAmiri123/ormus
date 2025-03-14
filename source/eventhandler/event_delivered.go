package eventhandler

import (
	"context"
	"fmt"

	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/pkg/channel"
	"github.com/rezaAmiri123/ormus/pkg/encoder"
	"github.com/rezaAmiri123/ormus/pkg/retry"
)

func (c Consumer) EventHasDeliveredToDestination(ctx context.Context, msg channel.Message) error {
	decodedEvent := encoder.DecodeProcessedEvent(string(msg.Body))
	err := c.eventService.EventHasDelivered(ctx, decodedEvent)
	if err != nil {
		logger.L().Error(fmt.Sprintf("err on change delivered status of events : %s", err.Error()))
		fn := func() error {
			return c.eventService.EventHasDelivered(ctx, decodedEvent)
		}
		e := retry.Do(fn, c.retryNumber)
		if e != nil {
			return e
		}
	}
	logger.L().Info(fmt.Sprintf("processed event event array: %v has been retrieved", decodedEvent.Events))
	err = msg.Ack()
	if err != nil {
		logger.L().Debug(fmt.Sprintf("ack failed for message : %s", err.Error()))
		return err
	}
	return nil
}
