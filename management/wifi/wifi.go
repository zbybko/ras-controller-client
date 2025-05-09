package wifi

import (
	"fmt"
	"ras/management/systemctl"
	"ras/management/wifi/hostapd"

	"github.com/charmbracelet/log"
)

var ErrNoWiFiService = fmt.Errorf("wi-fi service is not available on this system")

type WiFiInfo struct {
	Active     bool   `json:"active"`
	SSID       string `json:"ssid"`
	HiddenSSID bool   `json:"hidden_ssid"`
	Password   string `json:"password"`
	Security   string `json:"security"`
	Channel    int    `json:"channel"`
}

func Status() (*WiFiInfo, error) {

	active := systemctl.IsActive("hostapd")

	config, err := hostapd.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load hostapd config: %w", err)
	}

	ssid := config.GetSSID()
	hidden := config.GetHideSSID()
	password := config.GetPassword()
	security := config.GetSecurityType()
	channel := config.GetChannel()

	return &WiFiInfo{
		Active:     active,
		SSID:       ssid,
		HiddenSSID: hidden,
		Password:   password,
		Security:   security,
		Channel:    channel,
	}, nil
}

func Enable() error {
	err := systemctl.Enable(hostapd.Service)
	if err != nil {
		log.Errorf("Failed to enable Wi-Fi: %s", err)
		return err
	}

	return nil
}

func Disable() error {
	err := systemctl.Disable(hostapd.Service)
	if err != nil {
		log.Errorf("Failed to disable Wi-Fi: %s", err)
		return err
	}

	return nil
}
