package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/ormus/manager/service/authservice"
	"github.com/rezaAmiri123/ormus/manager/service/sourceservice"
	"github.com/rezaAmiri123/ormus/manager/service/userservice"
	"github.com/rezaAmiri123/ormus/manager/validator/sourcevalidator"
)

type Handler struct {
	sourceSvc   sourceservice.Service
	userSvc     userservice.Service
	validateSvc sourcevalidator.Validator
	authSvc     authservice.JWT
}

func New(sourceSvc sourceservice.Service,
	userSvc userservice.Service,
	validateSvc sourcevalidator.Validator,
	authSvc authservice.JWT,
) *Handler {
	return &Handler{
		sourceSvc:   sourceSvc,
		userSvc:     userSvc,
		validateSvc: validateSvc,
		authSvc:     authSvc,
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
