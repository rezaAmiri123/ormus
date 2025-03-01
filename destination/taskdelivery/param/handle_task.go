package param

import "github.com/rezaAmiri123/ormus/destination/entity/taskentity"

type DeliveryTaskResponse struct {
	FailedReason   *string
	Attempts       uint8
	DeliveryStatus taskentity.IntegrationDeliveryStatus
}
