package modems

import (
	"encoding/json"
	"ras/management/mmcli"
	"ras/utils"

	"github.com/charmbracelet/log"
)

func List() ([]string, error) {
	output, err := utils.Execute("mmcli", mmcli.ListModemsFlag, mmcli.JsonOutputFlag)
	if err != nil {
		log.Errorf("Failed get modems list: %s", err)
		return nil, err
	}
	modems := struct {
		List []string `json:"modem-list"`
	}{}
	err = json.Unmarshal(output, &modems)
	if err != nil {
		log.Errorf("Failed parse modem list from JSON: %s", err)
		return nil, err
	}
	return modems.List, nil
}

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
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--enable")
	return err
}

func (m *ModemInfo) GetSignal() (*ModemSignal, error) {
	output, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--signal-get", mmcli.JsonOutputFlag)
	if err != nil {
		log.Errorf("Failed get signal: %s", err)
		return nil, err
	}
	info := struct {
		Signal ModemSignal `json:"modem.signal"`
	}{}
	err = json.Unmarshal(output, &info)
	if err != nil {
		log.Errorf("Failed parse modem signal info from JSON: %s", err)
		return nil, err
	}
	return &info.Signal, nil
}

func (m *ModemInfo) SetPowerStateOff() error {
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-off")
	return err
}

func (m *ModemInfo) SetPowerStateOn() error {
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-on")
	return err
}
