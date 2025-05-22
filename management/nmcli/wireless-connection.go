package nmcli

import (
	"strconv"
)

const (
	OptionKeyWirelessSSID             = "802-11-wireless.ssid"
	OptionKeyWirelessHidden           = "802-11-wireless.hidden"
	OptionKeyWirelessChanel           = "802-11-wireless.chanel"
	OptionKeyWirelessMode             = "802-11-wireless.mode"
	OptionKeyWirelessSecurityPassword = "802-11-wireless-security.psk"
	OptionKeyWirelessSecurityKeyMgmt  = "802-11-wireless-security.key-mgmt" //Probably security mode
	OptionKeyWirelessSecurityProto    = "802-11-wireless-security.proto"
	OptionKeyWirelessSecurityGroup    = "802-11-wireless-security.group"
	OptionKeyWirelessSecurityPairwise = "802-11-wireless-security.pairwise"
)

func CreateWirelessConnection(
	deviceName string,
	connectionName string) (*WirelessConnection, error) {

	conn, err := createConnection(
		ConnectionTypeWIFI, deviceName, connectionName,
		[]string{"autoconnect", "yes", "ssid", connectionName},
	)
	if err != nil {
		return nil, err
	}
	wireless := WirelessConnection{conn}
	// TODO: idk what this consts mean, but they should be named
	wireless.setOption(OptionKeyWirelessSecurityKeyMgmt, "wpa-psk")
	wireless.setOption(OptionKeyWirelessSecurityProto, "rsn")
	wireless.setOption(OptionKeyWirelessSecurityGroup, "ccmp")
	wireless.setOption(OptionKeyWirelessSecurityPairwise, "ccmp")

	return &wireless, nil
}

type WirelessMode string

const (
	WirelessModeAccessPoint WirelessMode = "ap"
)

func (c *WirelessConnection) SetMode(mode WirelessMode) error {
	return c.setOption(OptionKeyWirelessMode, string(mode))
}

type WirelessBand = string

const (
	WirelessBandDefault WirelessBand = "bg"
)

func (c *WirelessConnection) SetBand(band WirelessBand) error {
	return c.setOption(OptionKeyWirelessMode, string(band))
}

func (c *WirelessConnection) GetSSID() string {
	return c.getOption(OptionKeyWirelessSSID)
}
func (c *WirelessConnection) SetSSID(ssid string) error {
	err := c.setOption(OptionKeyWirelessSSID, ssid)
	if err == nil {
		c.ensureOptionsParsed()
	}
	return err
}
func (c *WirelessConnection) GetChanel() int {
	opt := c.getOption(OptionKeyWirelessChanel)
	value, err := strconv.Atoi(opt)
	if err != nil {
		return 0
	}
	return value
}
func (c *WirelessConnection) SetChannel(chanel int) error {
	return c.setOption(OptionKeyWirelessChanel, strconv.Itoa(chanel))
}
func (c *WirelessConnection) GetPassword() string {
	return c.getOption(OptionKeyWirelessSecurityPassword)
}
func (c *WirelessConnection) SetPassword(password string) error {
	return c.setOption(OptionKeyWirelessSecurityPassword, password)
}

const (
	WirelessHiddenValue    = "yes"
	WirelessNotHiddenValue = "no"
)

func (c *WirelessConnection) IsHidden() bool {
	return c.getOption(OptionKeyWirelessHidden) == WirelessHiddenValue
}
func (c *WirelessConnection) SetHidden(hide bool) error {
	var value string
	if hide {
		value = WirelessHiddenValue
	} else {
		value = WirelessNotHiddenValue
	}
	err := c.setOption(OptionKeyWirelessHidden, value)
	if err == nil {
		c.ensureOptionsParsed()
	}
	return err
}
