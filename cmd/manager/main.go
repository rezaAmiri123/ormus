package main

import (
	"github.com/rezaAmiri123/ormus/config"
	"github.com/rezaAmiri123/ormus/manager/delivery/httpserver"
	"github.com/rezaAmiri123/ormus/manager/delivery/httpserver/userhandler"
	usermock "github.com/rezaAmiri123/ormus/manager/mock"
	"github.com/rezaAmiri123/ormus/manager/service/authservice"
	"github.com/rezaAmiri123/ormus/manager/service/userservice"
	"github.com/rezaAmiri123/ormus/manager/validator/uservalidator"
)

func main() {
	cfg := config.C()

	setupSvc := setupServices(cfg)

	server := httpserver.New(cfg, setupSvc)

	server.Server()
}

func setupServices(cfg config.Config) httpserver.SetupServicesResponse {
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig)
	unknownRepo := usermock.NewMockRepository(false)
	userSvc := userservice.New(jwt, unknownRepo)
	validateUserSvc := uservalidator.New(unknownRepo)

	userHand := userhandler.New(userSvc, validateUserSvc)

	return httpserver.SetupServicesResponse{
		UserHandler: userHand,
	}
}
