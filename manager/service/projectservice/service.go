package projectservice

import (
	"github.com/rezaAmiri123/ormus/manager/entity"
	"github.com/rezaAmiri123/ormus/pkg/channel/adapter/simple"
)

type Repository interface {
	Create(name, ID string) (entity.Project, error)
}

type Service struct {
	repo           Repository
	internalBroker *simple.ChannelAdapter
}

func New(repository Repository, internalBroker *simple.ChannelAdapter) *Service {
	return &Service{
		repo:           repository,
		internalBroker: internalBroker,
	}
}
