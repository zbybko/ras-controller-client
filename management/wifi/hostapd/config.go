package hostapd

import (
	"bufio"
	"fmt"
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
	HostapdConfigFile     = "/etc/hostapd/hostapd.conf"
	HostapdConfigFileMode = 0644

	Service = "hostapd.service"

	InterfaceKey = "interface"
	DriverKey    = "driver"
	SSIDKey      = "ssid"
	PasswordKey  = "wpa_passphrase"
	HideSSIDKey  = "ignore_broadcast_ssid"
	SecurityKey  = "wpa"
	ChannelKey   = "channel"

	MinPasswordLength = 8
	MaxPasswordLength = 63

	CommentSym = '#'
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
		if line == "" || line[0] == CommentSym {
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

func (c Config) GetChannel() int {
	if val, exists := c.conf[ChannelKey]; exists {
		val, err := strconv.Atoi(val)
		if err != nil {
			log.Errorf("Failed to parse integer value of hostapd channel: %s", err)
		}
		return val
	}
	log.Warn("No channel specified in hostapd config file", "configFile", HostapdConfigFile, "channelKey", ChannelKey)
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
	if err := Restart(); err != nil {
		log.Errorf("failed to restart hostapd: %s", err)
		return err
	}

	return os.WriteFile(HostapdConfigFile, []byte(output), HostapdConfigFileMode)
}

const (
	SSIDHidden    = "1"
	SSIDNotHidden = "0"
)

func (c Config) GetHideSSID() bool {
	return c.conf[HideSSIDKey] == SSIDHidden
}

func SetSSIDHidden(hidden bool) error {
	val := SSIDNotHidden
	if hidden {
		val = SSIDHidden
	}
	return updateConfig(HideSSIDKey, val)
}

func SetSSID(name string) error {
	return updateConfig(SSIDKey, name)
}

func SetPassword(password string) error {
	if len(password) < MinPasswordLength || len(password) > MaxPasswordLength {
		return fmt.Errorf("invalid passphrase length %d, expected %d..%d", len(password), MinPasswordLength, MaxPasswordLength)
	}
	return updateConfig(PasswordKey, password)
}

// Start Security

type SecurityType = string

const (
	SecurityTypeWPA2 SecurityType = "2"
	SecurityTypeWPA3 SecurityType = "3"
)

func (c Config) GetSecurityType() string {
	return c.conf[SecurityKey]
}

func SetSecurityType(st SecurityType) error {
	return updateConfig(SecurityKey, string(st))
}

// End Security

func SetChannel(channel int) error {
	return updateConfig(ChannelKey, strconv.Itoa(channel))
}

func (c Config) GetInterface() string {
	return c.conf[InterfaceKey]
}

func Enable() error {
	return systemctl.Enable(Service)
}
func Disable() error {
	return systemctl.Disable(Service)
}

func Restart() error {
	return systemctl.Restart(Service)
}
