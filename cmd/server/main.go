package main

import (
	"ras/config"
	"ras/server"
)

func main() {
	config.LoadConfigFile()
	config.SetupLogger()

	srv := server.New()

	srv.Run()
}
