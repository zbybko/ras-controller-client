package firewall

import (
	"ras/utils"
)

func Enable() error {
	_, err := utils.Execute("systemctl", "start", "firewalld")
	return err
}

func Disable() error {
	_, err := utils.Execute("systemctl", "stop", "firewalld")
	return err
}
