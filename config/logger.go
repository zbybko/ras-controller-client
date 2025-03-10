package config

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func setupLogLevel() {
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

}

func SetupLogger() {
	setupLogLevel()

	if prefix := viper.GetString("log.prefix"); prefix != "" {
		log.SetPrefix(prefix)
	}
}
