package eventhandler

import (
	"context"
	"fmt"

	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/pkg/channel"
	"github.com/rezaAmiri123/ormus/pkg/encoder"
	"github.com/rezaAmiri123/ormus/pkg/retry"
)

func (c Consumer) ProcessNewSourceEvent(ctx context.Context, msg channel.Message) error {
	decodedEvent := encoder.DecodeNewSourceEvent(string(msg.Body))

	logger.L().Info(fmt.Sprintf("project id : %s, write key: %s, owner id: %s: has been retrieved",
		decodedEvent.ProjectId, decodedEvent.WriteKey, decodedEvent.OwnerId))

	err := c.writeKeyService.CreateNewWriteKey(ctx, decodedEvent.OwnerId, decodedEvent.ProjectId, decodedEvent.WriteKey)
	if err != nil {
		logger.L().Error(fmt.Sprintf("err on creating writekey in redis : %s ", err.Error()))
		fn := func() error {
			return c.writeKeyService.CreateNewWriteKey(ctx, decodedEvent.OwnerId, decodedEvent.ProjectId, decodedEvent.WriteKey)
		}
		e := retry.Do(fn, c.retryNumber)
		if e != nil {
			return e
		}
	}

	logger.L().Debug("the message has been received")

	err = msg.Ack()
	if err != nil {
		logger.L().Debug(fmt.Sprintf("ack failed for message : %s", err.Error()))

		return err

	}

	return nil
}
