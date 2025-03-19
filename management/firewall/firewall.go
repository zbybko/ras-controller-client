package firewall

import (
	"os/exec"
	"ras/utils"
	"strings"

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

const ExitCodeInactive = 3
const ExitCodeNormal = 0

func Status() (*FirewallInfo, error) {
	output, err := utils.Execute("systemctl", "is-active", FirewallService)

	// Pass if error is 'exit code 3'
	// Exit code 3 is for inactive state of service
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() != ExitCodeInactive {
				log.Errorf("Failed get status of firewall, exit code isn't valid: %s", err)
			}
		} else {
			log.Errorf("Failed get status of firewall: %s", err)
			return nil, err

		}
	}
	strOutput := strings.TrimSpace(string(output))
	active := strOutput == FirewallStatusActive

	return &FirewallInfo{
		Active: active,
	}, nil
}

const FirewallService = "firewalld.service"

const FirewallStatusActive = "active"
const FirewallStatusInactive = "inactive"
