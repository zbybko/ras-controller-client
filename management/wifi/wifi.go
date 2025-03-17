package wifi

import (
	"fmt"
	"ras/utils"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

var ErrNoWiFiService = fmt.Errorf("Wi-Fi service is not available on this system")

type RealWiFiManager struct{}

type WiFiInfo struct {
	Active     bool   `json:"active"`
	SSID       string `json:"ssid"`
	HiddenSSID bool   `json:"hidden_ssid"`
	Password   string `json:"password"`
	Security   string `json:"security"`
	Channel    int    `json:"channel"`
}

// Проверка статуса Wi-Fi
func (r *RealWiFiManager) Status() (*WiFiInfo, error) {
	if !r.ServiceExists() {
		return nil, ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "--status")
	if err != nil {
		log.Errorf("Failed to get Wi-Fi status: %s", err)
		return nil, err
	}

	strOutput := strings.TrimSpace(string(output))

	active := false
	if match, _ := regexp.MatchString("enabled", strOutput); match {
		active = true
	} else if match, _ := regexp.MatchString("disabled", strOutput); match {
		active = false
	}

	return &WiFiInfo{
		Active: active,
	}, nil
}

// Включить Wi-Fi
func (r *RealWiFiManager) Enable() error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "--enable")
	if err != nil {
		log.Errorf("Failed to enable Wi-Fi: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("enabled", string(output)); !match {
		return fmt.Errorf("failed to enable Wi-Fi: %s", output)
	}

	log.Info("Wi-Fi enabled successfully")
	return nil
}

// Выключить Wi-Fi
func (r *RealWiFiManager) Disable() error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "--disable")
	if err != nil {
		log.Errorf("Failed to disable Wi-Fi: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("disabled", string(output)); !match {
		return fmt.Errorf("failed to disable Wi-Fi: %s", output)
	}

	log.Info("Wi-Fi disabled successfully")
	return nil
}

// Проверка доступности Wi-Fi сервиса
func (r *RealWiFiManager) ServiceExists() bool {
	_, err := utils.Execute("mmcli", "--version")
	return err == nil
}

// Установить скрытие SSID
func (r *RealWiFiManager) SetSSIDHidden(hidden bool) error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}
	state := "no"
	if hidden {
		state = "yes"
	}

	output, err := utils.Execute("mmcli", "-m", "0", "modem", "set-ssid", "Hotspot", state)
	if err != nil {
		log.Errorf("Failed to set SSID hidden state: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("success", string(output)); !match {
		return fmt.Errorf("failed to set SSID hidden state: %s", output)
	}

	log.Info("SSID hidden state set successfully")
	return nil
}

// Изменить SSID
func (r *RealWiFiManager) SetSSID(name string) error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "modem", "set-ssid", name)
	if err != nil {
		log.Errorf("Failed to set SSID: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("success", string(output)); !match {
		return fmt.Errorf("failed to set SSID: %s", output)
	}

	log.Info("SSID set successfully")
	return nil
}

// Установить пароль
func (r *RealWiFiManager) SetPassword(password string) error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "modem", "set-password", password)
	if err != nil {
		log.Errorf("Failed to set password: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("success", string(output)); !match {
		return fmt.Errorf("failed to set password: %s", output)
	}

	log.Info("Password set successfully")
	return nil
}

// Установить тип безопасности
func (r *RealWiFiManager) SetSecurityType(wpa3 bool) error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}
	keyMgmt := "wpa-psk"
	if wpa3 {
		keyMgmt = "sae"
	}

	output, err := utils.Execute("mmcli", "-m", "0", "modem", "set-encryption", keyMgmt)
	if err != nil {
		log.Errorf("Failed to set encryption type: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("success", string(output)); !match {
		return fmt.Errorf("failed to set encryption type: %s", output)
	}

	log.Info("Encryption type set successfully")
	return nil
}

// Установить канал
func (r *RealWiFiManager) SetChannel(channel int) error {
	if !r.ServiceExists() {
		return ErrNoWiFiService
	}

	output, err := utils.Execute("mmcli", "-m", "0", "modem", "set-channel", fmt.Sprintf("%d", channel))
	if err != nil {
		log.Errorf("Failed to set channel: %s", err)
		return err
	}

	if match, _ := regexp.MatchString("success", string(output)); !match {
		return fmt.Errorf("failed to set channel: %s", output)
	}

	log.Info("Channel set successfully")
	return nil
}
