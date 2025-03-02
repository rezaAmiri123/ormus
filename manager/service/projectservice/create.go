package projectservice

import (
	"github.com/rezaAmiri123/ormus/param"
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
)

func (s Service) Create(req param.CreateProjectRequest) (*param.CreateProjectResponse, error) {
	const op = "projectservice.Create"

	newProject, err := s.repo.Create(req.Name, req.UserID)
	if err != nil {
		return nil, richerror.New(op).WithWrappedError(err).WhitMessage(errmsg.ErrSomeThingWentWrong)
	}

	return &param.CreateProjectResponse{
		ID:   newProject.ID,
		Name: newProject.Name,
	}, nil
}
