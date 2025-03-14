package main

import (
	"ras/config"
	"ras/server"

	"github.com/charmbracelet/log"
)

func main() {
	config.LoadConfigFile()
	config.SetupLogger()

	srv := server.New()

	log.Debug("Starting server")
	if err := srv.Run(server.Address()); err != nil {
		log.Fatalf("Failed start server: %s", err)
	}
}
