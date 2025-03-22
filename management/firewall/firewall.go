package firewall

import (
	"fmt"
	"os/exec"
	"ras/management/systemctl"
	"ras/utils"
	"strings"

	"github.com/charmbracelet/log"
)

var NoServiceErr = fmt.Errorf("there is not firewall service in system")

func Enable() error {
	if !ServiceExists() {
		return NoServiceErr
	}
	return systemctl.Enable(FirewallService)
}

func Disable() error {
	if !ServiceExists() {
		return NoServiceErr
	}
	return systemctl.Disable(FirewallService)
}

type FirewallInfo struct {
	Active bool `json:"active"`
}

const ExitCodeInactive = 3
const ExitCodeNormal = 0

func ServiceExists() bool {
	return systemctl.ServiceExists(FirewallService)
}

func Status() (*FirewallInfo, error) {
	if !ServiceExists() {
		return nil, NoServiceErr
	}
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
