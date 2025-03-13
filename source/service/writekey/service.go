package writekey

import (
	"context"

	proto_source "github.com/rezaAmiri123/ormus/contract/go/source"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
	"github.com/rezaAmiri123/ormus/source"
)

type Repository interface {
	CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error
	GetWriteKey(ctx context.Context, ownerID, projectID string) (*proto_source.NewSourceEvent, error)
}

type Service struct {
	writeKeyRepo Repository
	config       source.Config
}

func New(writeKeyRepo Repository, config source.Config) Service {
	return Service{
		writeKeyRepo: writeKeyRepo,
		config:       config,
	}
}

func (s Service) CreateNewWriteKey(ctx context.Context, ownerID, projectID, writeKey string) error {
	err := s.writeKeyRepo.CreateNewWriteKey(ctx, &proto_source.NewSourceEvent{
		ProjectId: projectID,
		OwnerId:   ownerID,
		WriteKey:  writeKey,
	}, s.config.WriteKeyRedisExpiration)
	if err != nil {
		return richerror.New("source.service").WithWrappedError(err)
	}

	return nil
}
