package config

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var logLevel = log.InfoLevel

func setupLogLevel() {
	levelString := viper.GetString("log.level")
	switch levelString {
	case "info":
	case "debug":
		logLevel = log.DebugLevel
	case "warning":
		logLevel = log.WarnLevel
	}
	if viper.GetBool("debug") {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
}

func GetLogger(prefix string) *log.Logger {
	l := log.Default().WithPrefix(prefix)
	l.SetLevel(logLevel)
	return l
}

func SetupLogger() {
	setupLogLevel()

	if prefix := viper.GetString("log.prefix"); prefix != "" {
		log.SetPrefix(prefix)
	}
}
