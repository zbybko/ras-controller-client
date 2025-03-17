package sim

import (
	"encoding/json"
	"ras/management/mmcli"
	"ras/utils"

	"github.com/charmbracelet/log"
)

func Get(sim string) (*SimInfo, error) {
	output, err := utils.Execute("mmcli", mmcli.SimFlag(sim), mmcli.JsonOutputFlag)

	if err != nil {
		log.Errorf("Failed get sim: %s", err)
		return nil, err
	}

	info := struct {
		Sim SimInfo `json:"sim"`
	}{}
	err = json.Unmarshal(output, &info)
	if err != nil {
		log.Errorf("Failed parse sim info from JSON: %s", err)
		return nil, err
	}
	return &info.Sim, nil
}
