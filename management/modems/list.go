package modems

import (
	"encoding/json"
	"fmt"
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
		log.Warnf("Error occurred while getting signal, and was ignored: %s", err)
		log.Warn("This error should not be ignored! Call Katy248 to fix it ASAP")
		// log.Errorf("Failed get signal: %s", err)
		// return nil, err
	}
	info := struct {
		Modem struct {
			Signal ModemSignal `json:"signal"`
		} `json:"modem"`
	}{}
	err = json.Unmarshal(output, &info)
	if err != nil {
		log.Errorf("Failed parse modem signal info from JSON: %s", err)
		return nil, err
	}
	log.Debugf("Modem signal json string: %s", string(output))
	log.Debugf("Modem signal: %+v", info.Modem.Signal)
	return &info.Modem.Signal, nil
}

func (m *ModemInfo) SetPowerStateOff() error {
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-off")
	return err
}

func (m *ModemInfo) SetPowerStateOn() error {
	_, err := utils.Execute("mmcli", mmcli.ModemFlag(m.DBusPath), "--set-power-state-on")
	return err
}
func (m *ModemInfo) GetBearer() (*BearerInfo, error) {
	if len(m.Generic.Bearers) == 0 {
		return nil, fmt.Errorf("no bearers")
	}
	_, err := utils.Execute("mmcli", mmcli.BearerFlag(m.Generic.Bearers[0]), mmcli.JsonOutputFlag)
	var info struct {
		Bearer BearerInfo `json:"bearer"`
	}
	return &info.Bearer, err
}
