package eventvalidator

import (
	"context"

	proto_source "github.com/rezaAmiri123/ormus/contract/go/source"
	"github.com/rezaAmiri123/ormus/source"
)

type Repository interface {
	IsWriteKeyValid(ctx context.Context, writeKey string, expirationTime uint) (bool, error)
	CreateNewWriteKey(ctx context.Context, writeKey *proto_source.NewSourceEvent, expirationTime uint) error
}

type Validator struct {
	repo   Repository
	config source.Config
}

func New(repo Repository, cfg source.Config) Validator {
	return Validator{
		repo:   repo,
		config: cfg,
	}
}
