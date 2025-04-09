package diagnostics

import (
	"ras/utils"

	"github.com/charmbracelet/log"
)

var defaultPingAddress = ""

func GetDefaultPingAddress() string {
	if defaultPingAddress == "" {
		setupDiagnostics()
	}
	return defaultPingAddress
}

func Ping(addr string) (string, error) {
	if addr == "" {
		log.Debugf("Changing empty ping address to default")
		addr = GetDefaultPingAddress()
	}

	output, err := utils.Execute("ping", addr, "-c", "5")
	return string(output), err
}
