package firewall

import (
	"fmt"
	"ras/management/systemctl"
	"ras/utils"

	"github.com/charmbracelet/log"
)

const FirewallService = "firewalld.service"

var (
	ErrNoService        = fmt.Errorf("there is not firewall service in system")
	ErrServiceNotActive = fmt.Errorf("firewall service is not active (maybe disabled)")
)

func Enable() error {
	if !ServiceExists() {
		return ErrNoService
	}
	return systemctl.Enable(FirewallService)
}

func Disable() error {
	if !ServiceExists() {
		return ErrNoService
	}
	return systemctl.Disable(FirewallService)
}

type FirewallInfo struct {
	Active        bool `json:"active"`
	ServiceExists bool `json:"serviceExists"`
}

func ServiceExists() bool {
	return systemctl.ServiceExists(FirewallService)
}

func IsActive() bool {
	return systemctl.IsActive(FirewallService)
}

func Status() (*FirewallInfo, error) {
	if !ServiceExists() {
		return &FirewallInfo{}, ErrNoService
	}

	return &FirewallInfo{
		ServiceExists: ServiceExists(),
		Active:        IsActive(),
	}, nil
}

func AddService(serviceName string) error {
	if !IsActive() {
		return ErrServiceNotActive
	}
	if err := utils.ExecuteErr("firewall-cmd", "--permanent", fmt.Sprintf("--add-service=%s", serviceName)); err != nil {
		log.Errorf("Failed to add service '%s' : %s", serviceName, err)
		return err
	}

	return Reload()
}

func Reload() error {
	if !IsActive() {
		return ErrServiceNotActive
	}
	return utils.ExecuteErr("firewall-cmd", "--reload")
}

func Restart() error {
	if !ServiceExists() {
		return ErrNoService
	}

	if !IsActive() {
		return ErrServiceNotActive
	}
	return systemctl.Restart(FirewallService)
}
