package nmcli

import (
	"ras/utils"
	"strings"
)

const (
	HardwareAddressField = "GENERAL.HWADDR"
)

func GetHardwareAddress(deviceName string) (string, error) {
	output, err := utils.Execute("nmcli", terseFlag, getFields(HardwareAddressField), "device", "show", deviceName)
	return strings.ReplaceAll(strings.TrimSpace(string(output)), `\`, ""), err
}
