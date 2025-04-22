package config

import (
	"github.com/charmbracelet/log"

	"github.com/spf13/viper"
)

const defaultConfigName = "config"

func init() {
	log.Debug("Initializing configurations")
	viper.SetEnvPrefix("RAS")
	viper.AutomaticEnv()

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	viper.SetDefault("config", defaultConfigName)
	viper.SetDefault("use_mock", false)
	configName := viper.GetString("config")
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
}

func LoadConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debugf("Config file not found")
			// Config file not found; ignore error if desired
		} else {
			log.Fatalf("Failed to load config file: %s", err)
		}
	}
}
func Save() {
	err := viper.WriteConfig()
	if err != nil {
		log.Errorf("Failed to save config file: %s", err)
	}
	log.Debug("Saved config file")
}
