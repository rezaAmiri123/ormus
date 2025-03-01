package fakedeliveryhandler

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/rezaAmiri123/ormus/destination/entity/taskentity"
	"github.com/rezaAmiri123/ormus/destination/taskdelivery/param"
	"github.com/rezaAmiri123/ormus/event"
)

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

const fakeProcessingTimeSecond = 2

func (h FakeHandler) Handle(t taskentity.Task, _ event.ProcessedEvent) (param.DeliveryTaskResponse, error) {
	time.Sleep(fakeProcessingTimeSecond * time.Second)

	slog.Info(fmt.Sprintf("Task [%s] handled successfully!", t.ID))

	res := param.DeliveryTaskResponse{
		Attempts:       1,
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}

	return res, nil
}
