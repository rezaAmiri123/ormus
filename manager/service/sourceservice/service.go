package sourceservice

import (
	"github.com/rezaAmiri123/ormus/manager/entity"
	"github.com/rezaAmiri123/ormus/manager/managerparam"
)

type SourceRepo interface {
	InsertSource(source *entity.Source) (*managerparam.AddSourceResponse, error)
	UpdateSource(id string, source *entity.Source) (*managerparam.UpdateSourceResponse, error)
	DeleteSource(id, userID string) error
	GetUserSourceByID(ownerID, id string) (*entity.Source, error)
	IsSourceAlreadyCreatedByName(name string) (bool, error)
}

type Service struct {
	repo SourceRepo
}

func New(repo SourceRepo) *Service {
	return &Service{
		repo: repo,
	}
}
