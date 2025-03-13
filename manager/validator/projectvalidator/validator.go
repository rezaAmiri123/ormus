package projectvalidator

import "github.com/rezaAmiri123/ormus/manager/entity"

type Validator struct {
	repo Repository
}

type Repository interface {
	GetWithID(id string) (entity.Project, error)
}

func New(repo Repository) Validator {
	return Validator{
		repo: repo,
	}
}
