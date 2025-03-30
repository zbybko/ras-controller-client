package modems

import (
	"ras/management/mmcli"
	"ras/utils"

	"github.com/charmbracelet/log"
)

func (m *ModemInfo) SetPowerStateOff() error {
	log.Debugf("Set power state off for modem %s", m.DBusPath)
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-off")
	return err
}

func (m *ModemInfo) SetPowerStateOn() error {
	log.Debugf("Set power state on for modem %s", m.DBusPath)
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-on")
	return err
}
