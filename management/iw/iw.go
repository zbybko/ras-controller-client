package iw

import (
	"errors"
	"ras/utils"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type ConnectedDevice struct {
	MAC       string
	Interface string
	TxBitrate string
	RxBitrate string
}

func GetConnectedDevices() ([]ConnectedDevice, error) {
	output, err := utils.Execute("iw", "dev", viper.GetString("wifi.default_interface"), "station", "dump")
	if err != nil {
		return nil, err
	}
	return parseConnectedDevices(string(output))
}

func parseConnectedDevices(output string) ([]ConnectedDevice, error) {
	blocks := strings.Split(output, "Station")

	devices := []ConnectedDevice{}
	for _, block := range blocks {
		d, err := parseDevice(block)
		if err != nil {
			log.Warnf("Failed parsing device block: %s", err)
			continue
		}
		devices = append(devices, *d)
	}
	return devices, nil

}

var (
	ErrBadDeviceBlock = errors.New("bad device block")
	ErrBadFirstLine   = errors.New("bad first line")
)

func parseDevice(block string) (*ConnectedDevice, error) {
	lines := strings.Split(block, "\n")
	if len(lines) < 1 {
		return nil, ErrBadDeviceBlock
	}

	device, err := parseFirstLine(lines[0])
	if err != nil {
		return nil, err
	}

	dict := parseLines(lines[1:])
	device.RxBitrate = dict["rx bitrate"]
	device.TxBitrate = dict["tx bitrate"]
	return &device, nil
}

func parseLines(lines []string) map[string]string {
	linesMap := make(map[string]string)
	for _, line := range lines {
		words := strings.Split(line, ":")
		if len(words) < 2 {
			continue
		}
		key := strings.TrimSpace(words[0])
		value := strings.TrimSpace(words[1])
		linesMap[key] = value
	}
	return linesMap
}

// Line should not contain 'Station'
func parseFirstLine(line string) (ConnectedDevice, error) {
	device := ConnectedDevice{}

	words := strings.Split(line, " ")
	if len(words) < 3 {
		return device, ErrBadFirstLine
	}
	device.MAC = words[0]
	device.Interface = strings.Replace(words[2], ")", "", 1)
	return device, nil
}
