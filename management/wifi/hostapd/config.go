package hostapd

import (
	"bufio"
	"os"
	"strings"
)

const (
	HostapdConfigFile = "/etc/hostapd/hostapd.conf"

	InterfaceKey = "interface"
	DriverKey    = "driver"
	SSIDKey      = "ssid"
	PasswordKey  = "wps_passphrase"
	HideSSIDKey  = "ignore_broadcast_ssid"
)

func New() (*Config, error) {
	configFile, err := os.Open(HostapdConfigFile)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	conf := Config{
		conf: make(map[string]string),
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		keyVal := strings.Split(line, "=")
		if len(keyVal) != 2 {
			continue
		}
		conf.conf[keyVal[0]] = keyVal[1]
	}
	return &conf, nil
}

type Config struct {
	conf map[string]string
}

func (c Config) GetSSID() string {
	return c.conf[SSIDKey]
}

func (c Config) GetPassword() string {
	return c.conf[PasswordKey]
}

const SSIDHidden = "1"

func (c Config) GetHideSSID() bool {
	return c.conf[HideSSIDKey] == SSIDHidden
}
