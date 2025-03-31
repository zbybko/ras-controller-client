package dhcp

import (
	"errors"
	"ras/management/systemctl"
	"ras/utils"
)

type DhcpStatus struct {
	Enabled bool `json:"enabled"`
}

const DhcpService = "dhcpd.service"

func Status() *DhcpStatus {
	return &DhcpStatus{
		Enabled: systemctl.IsActive(DhcpService),
	}
}

func Enable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	return errors.Join(
		systemctl.Enable(DhcpService),
		utils.ExecuteErr("firewall-cmd", "--permament", "--add-service=dhcp"),
		utils.ExecuteErr("firewall-cmd", "--reload"),
	)
}
func Disable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	return systemctl.Disable(DhcpService)
}
