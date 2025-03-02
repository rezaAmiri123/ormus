package uservalidator

import (
	"github.com/rezaAmiri123/ormus/manager/service/userservice"
)

type Validator struct {
	repo userservice.Repository
}

func New(repo userservice.Repository) Validator {
	return Validator{repo: repo}
}
