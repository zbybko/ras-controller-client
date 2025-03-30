package modems

import (
	"encoding/json"
	"ras/management/mmcli"
	"ras/utils"

	"github.com/charmbracelet/log"
)

func Get(modem string) (*ModemInfo, error) {
	output, err := utils.Execute("mmcli", mmcli.ModemFlag(modem), mmcli.JsonOutputFlag)
	if err != nil {
		log.Errorf("Failed get modem info: %s", err)
		return nil, err
	}
	info := struct {
		Modem ModemInfo `json:"modem"`
	}{}
	err = json.Unmarshal(output, &info)
	if err != nil {
		log.Errorf("Failed parse modem info from JSON: %s", err)
		return nil, err
	}
	return &info.Modem, nil
}

func (m *ModemInfo) Disable() error {
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--disable")
	return err
}

func (m *ModemInfo) Enable() error {
	m.SetPowerStateOn() // ensures that it turned on
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--enable")
	return err
}
