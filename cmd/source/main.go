package main

import (
	"github.com/rezaAmiri123/ormus/config"
	"github.com/rezaAmiri123/ormus/source/delivery/httpserver"
	"github.com/rezaAmiri123/ormus/source/delivery/httpserver/statushandler"
)

func main() {
	handlers := []httpserver.Handler{
		statushandler.New(),
	}

	httpServer := httpserver.New(config.C().Source, handlers)

	httpServer.Serve()
}
