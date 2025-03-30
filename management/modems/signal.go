package modems

import (
	"encoding/json"
	"ras/management/mmcli"
	"ras/utils"

	"github.com/charmbracelet/log"
)

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
