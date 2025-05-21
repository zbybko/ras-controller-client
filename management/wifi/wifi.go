package wifi

import (
	"fmt"
	"ras/management/systemctl"
	"ras/management/wifi/hostapd"

	"github.com/charmbracelet/log"
)

var ErrNoWiFiService = fmt.Errorf("wi-fi service is not available on this system")

var config *hostapd.Config

type WiFiInfo struct {
	Active     bool   `json:"active"`
	SSID       string `json:"ssid"`
	HiddenSSID bool   `json:"hidden_ssid"`
	Password   string `json:"password"`
	Security   string `json:"security"`
	Channel    int    `json:"channel"`
}

func Status() (*WiFiInfo, error) {
	var err error
	active := systemctl.IsActive("hostapd")

	if config == nil {

	}

	config, err = hostapd.New()

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
	err := hostapd.Enable()
	if err != nil {
		log.Errorf("Failed to enable Wi-Fi: %s", err)
		return err
	}

	return nil
}

func Disable() error {
	err := hostapd.Disable()
	if err != nil {
		log.Errorf("Failed to disable Wi-Fi: %s", err)
		return err
	}

	return nil
}
