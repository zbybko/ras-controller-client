package modems

import (
	"encoding/json"
	"fmt"
	"ras/utils"

	"github.com/charmbracelet/log"
)

const MMCLIJsonOutputFlag = "--output-json"
const MMCLIListModemsFlag = "--list-modems"

func GetList() ([]string, error) {
	output, err := utils.Execute("mmcli", MMCLIListModemsFlag, MMCLIJsonOutputFlag)
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

func modemFlag(modem string) string {
	return fmt.Sprintf("--modem='%s'", modem)
}

func GetInfo(modem string) (*ModemInfo, error) {
	output, err := utils.Execute("mmcli", modemFlag(modem), MMCLIJsonOutputFlag)
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
	_, err := utils.Execute("mmcli", modemFlag(m.DBusPath), "--disable")
	return err
}
func (m *ModemInfo) Enable() error {
	_, err := utils.Execute("mmcli", modemFlag(m.DBusPath), "--enable")
	return err
}

func (m *ModemInfo) GetSignal() (*ModemSignal, error) {
	output, err := utils.Execute("mmcli", modemFlag(m.DBusPath), "--get-signal", MMCLIJsonOutputFlag)
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
