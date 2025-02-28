package sourceservice

import (
	"github.com/rezaAmiri123/ormus/manager/entity"
	"github.com/rezaAmiri123/ormus/manager/param"
	writekey "github.com/rezaAmiri123/ormus/pkg/write_key"
)

func (s Service) CreateSource(req *param.AddSourceRequest, ownerID string) (*param.AddSourceResponse, error) {
	w, err := writekey.GenerateNewWriteKey()
	if err != nil {
		return nil, err
	}

	source := &entity.Source{
		ID:          "", // TODO  uuid ulid ?
		WriteKey:    entity.WriteKey(w),
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     ownerID,
		ProjectID:   req.ProjectID,
	}

	response, err := s.repo.InsertSource(source)
	if err != nil {
		return nil, err
	}

	return response, nil
}
