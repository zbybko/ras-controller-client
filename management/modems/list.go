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
