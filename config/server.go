package config

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const ServerModeKey = "server.mode"

var allowedModes = []string{gin.DebugMode, gin.ReleaseMode, gin.TestMode}

func init() {
	viper.SetDefault(ServerModeKey, gin.DebugMode)
}

func applyServerMode(mode string) error {
	if !slices.Contains(allowedModes, mode) {

		log.Errorf("Invalid server mode '%s'. Valid modes: %v", mode, allowedModes)
		return fmt.Errorf("invalid server mode '%s'", mode)
	}
	gin.SetMode(mode)
	if gin.Mode() != mode {
		log.Errorf("Failed to set server mode to '%s', probably internal gin error", mode)
		return fmt.Errorf("mode is not set up, internal error")
	}
	return nil
}

func SetServerMode(mode string) {
	if err := applyServerMode(mode); err != nil {
		log.Errorf("Failed to set server mode to '%s': %s", mode, err)
		return
	}
	viper.Set(ServerModeKey, mode)
	Save()
}

func GetServerMode() string {
	if viper.GetBool("debug") {
		return gin.DebugMode
	}
	return viper.GetString(ServerModeKey)
}

func SetupServerMode() {
	mode := GetServerMode()
	if err := applyServerMode(mode); err != nil {
		log.Errorf("Failed to set server mode to '%s': %s", mode, err)
		return
	}
	log.Debugf("Server mode set to '%s'", GetServerMode())
}
