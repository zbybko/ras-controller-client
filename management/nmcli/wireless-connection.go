package nmcli

import (
	"strconv"
)

const (
	OptionKeyWirelessSSID     = "802-11-wireless.ssid"
	OptionKeyWirelessHidden   = "802-11-wireless.hidden"
	OptionKeyWirelessChanel   = "802-11-wireless.chanel"
	OptionKeyWirelessPassword = "wifi-sec.psk"
	OptionKeyWirelessMode     = "802-11-wireless.mode"
	OptionKeyWirelessKeyMgmt  = "wifi-sec.key-mgmt" //Probably security mode
)

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
	return c.getOption(OptionKeyWirelessPassword)
}
func (c *WirelessConnection) SetPassword(password string) error {
	return c.setOption(OptionKeyWirelessPassword, password)
}

const (
	WirelessHiddenValue    = "no"
	WirelessNotHiddenValue = "yes"
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
