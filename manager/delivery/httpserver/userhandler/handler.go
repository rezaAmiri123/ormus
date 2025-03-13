package userhandler

import (
	"github.com/rezaAmiri123/ormus/manager/service/projectservice"
	"github.com/rezaAmiri123/ormus/manager/service/userservice"
)

type Handler struct {
	// TODO - add configurations
	userSvc    userservice.Service
	projectSvc projectservice.Service
}

func New(userSvc userservice.Service, projectSvc projectservice.Service) Handler {
	return Handler{
		userSvc:    userSvc,
		projectSvc: projectSvc,
	}
}
