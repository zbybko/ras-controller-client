package firewall

import (
	"ras/utils"

	"github.com/charmbracelet/log"
)

func Enable() error {
	_, err := utils.Execute("systemctl", "enable", "--now", FirewallService)
	return err
}

func Disable() error {
	_, err := utils.Execute("systemctl", "disable", "--now", FirewallService)
	return err
}

type FirewallInfo struct {
	Active bool `json:"active"`
}

func Status() (*FirewallInfo, error) {
	output, err := utils.Execute("systemctl", "is-active", FirewallService)
	if err == nil || err.Error() == "exit status 3" {
		active := string(output) == FirewallStatusActive
		return &FirewallInfo{
			Active: active,
		}, nil

	} else {
		log.Errorf("Failed get status of firewall: %s", err)
		return nil, err
	}
}

const FirewallService = "firewalld.service"

const FirewallStatusActive = "active"
const FirewallStatusInactive = "inactive"
