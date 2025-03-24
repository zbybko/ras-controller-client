package firewall

import (
	"fmt"
	"ras/management/systemctl"
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

func ServiceExists() bool {
	return systemctl.ServiceExists(FirewallService)
}

func Status() (*FirewallInfo, error) {
	if !ServiceExists() {
		return nil, NoServiceErr
	}
	isActive := systemctl.IsActive(FirewallService)

	return &FirewallInfo{
		Active: isActive,
	}, nil
}

const FirewallService = "firewalld.service"
