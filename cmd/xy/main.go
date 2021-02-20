package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"xy/pkg/cfg"
	"xy/pkg/env"
	"xy/pkg/server"
)

func main() {
	fmt.Println("xy")

	// generate config from env and cmd flags
	config := cfg.Config{}
	env.LoadEnvironment(&config)
	env.ParseFlags(&config)

	cfg.LoadDatabase(&config)

	// create router & bind routes
	router := mux.NewRouter()
	server.RegisterRoutes(router)

	fmt.Println("Server is starting...")

	// await server start
	go server.NotifyStarted(config)

	// start server
	server.Start(config, router)
}
