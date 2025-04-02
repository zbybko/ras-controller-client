package hostapd

import (
	"bufio"
	"os"
	"ras/management/systemctl"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Config struct {
	conf map[string]string
}

const (
	HostapdConfigFile = "/etc/hostapd/hostapd.conf"

	Service = "hostapd.service"

	InterfaceKey = "interface"
	DriverKey    = "driver"
	SSIDKey      = "ssid"
	PasswordKey  = "wps_passphrase"
	HideSSIDKey  = "ignore_broadcast_ssid"
	SecurityKey  = "wpa"
	ChannelKey   = "channel"
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

func (c Config) GetSecurityType() string {
	return c.conf[SecurityKey]
}

func (c Config) GetChannel() int {
	if val, exists := c.conf[ChannelKey]; exists {
		val, err := strconv.Atoi(val)
		if err != nil {
			log.Errorf("Failed to parseint: %s", err)
		}
		return val
	}
	return 0
}

func updateConfig(key, value string) error {
	input, err := os.ReadFile(HostapdConfigFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = key + "=" + value
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, key+"="+value)
	}

	output := strings.Join(lines, "\n")
	if err := RestartHostapd(); err != nil {
		log.Errorf("failed to restart hostapd: %s", err)
		return err
	}

	return os.WriteFile(HostapdConfigFile, []byte(output), 0644)
}

func SetSSIDHidden(hidden bool) error {
	val := "0"
	if hidden {
		val = SSIDHidden
	}
	return updateConfig(HideSSIDKey, val)
}

func SetSSID(name string) error {
	return updateConfig(SSIDKey, name)
}

func SetPassword(password string) error {
	return updateConfig(PasswordKey, password)
}

func SetSecurityType(wpa3 bool) error {
	val := "2"
	if wpa3 {
		val = "3"
	}
	return updateConfig(SecurityKey, val)
}

func SetChannel(channel int) error {
	return updateConfig(ChannelKey, strconv.Itoa(channel))
}

func RestartHostapd() error {
	return systemctl.Restart(Service)
}
