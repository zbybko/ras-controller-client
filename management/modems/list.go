package modems

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"ras/utils"

	"github.com/charmbracelet/log"
)

const MMCLIJsonOutputFlag = "--output-json"
const MMCLIListModemsFlag = "--list-modems"

func GetList() ([]string, error) {
	cmd := exec.Command("mmcli", MMCLIListModemsFlag, MMCLIJsonOutputFlag)
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("Failed execute command `%s`: %s", cmd.String(), err)
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
	cmd := exec.Command("mmcli", modemFlag(modem), MMCLIJsonOutputFlag)
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("Failed execute command `%s`: %s", cmd.String(), err)
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
	cmd := exec.Command("mmcli", modemFlag(m.DBusPath), "--disable")
	return cmd.Run()
}
func (m *ModemInfo) Enable() error {
	cmd := exec.Command("mmcli", modemFlag(m.DBusPath), "--enable")
	return cmd.Run()
}

func (m *ModemInfo) GetSignal() (*ModemSignal, error) {
	output, err := utils.Execute("mmcli", modemFlag(m.DBusPath), "--get-signal", MMCLIJsonOutputFlag)
	if err != nil {
		log.Errorf("Failed get list of modems: %s", err)
		return nil, err
	}
	info := struct {
		Signal ModemSignal `json:"modem.signal"`
	}{}
	err = json.Unmarshal([]byte(output), &info)
	if err != nil {
		log.Errorf("Failed parse modem signal info from JSON: %s", err)
		return nil, err
	}
	return &info.Signal, nil
}
