package main

import (
	"log"
	"ras/config"
	"ras/server"
)

func main() {
	config.LoadConfigFile()
	config.SetupLogger()

	srv := server.New()

	if err := srv.Run(server.Address()); err != nil {
		log.Fatalf("Failed start server: %s", err)
	}
}
