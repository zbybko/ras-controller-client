// Documentation for wireless nmcli:
//
// - https://www.networkmanager.dev/docs/api/latest/settings-802-11-wireless.html
//
// - https://www.networkmanager.dev/docs/api/latest/settings-802-11-wireless-security.html
package nmcli

import (
	"strconv"
)

const (
	OptionKeyWirelessSSID                  = "802-11-wireless.ssid"
	OptionKeyWirelessHidden                = "802-11-wireless.hidden"
	OptionKeyWirelessChanel                = "802-11-wireless.chanel"
	OptionKeyWirelessMode                  = "802-11-wireless.mode"
	OptionKeyWirelessBand                  = "802-11-wireless.band"
	OptionKeyWirelessSecurityPassword      = "802-11-wireless-security.psk"
	OptionKeyWirelessSecurityKeyManagement = "802-11-wireless-security.key-mgmt"
	OptionKeyWirelessSecurityProto         = "802-11-wireless-security.proto"
	OptionKeyWirelessSecurityGroup         = "802-11-wireless-security.group"
	OptionKeyWirelessSecurityPairwise      = "802-11-wireless-security.pairwise"
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
	// TODO: This options should be moved to own functions
	wireless.setOption(OptionKeyWirelessSecurityKeyManagement, KeyManagementWPA2_3Personal)
	wireless.setOption(OptionKeyWirelessSecurityProto, ProtoAllowWPA2RSN)
	wireless.setOption(OptionKeyWirelessSecurityGroup, EncryptionAlgCcmp)
	wireless.setOption(OptionKeyWirelessSecurityPairwise, EncryptionAlgCcmp)

	return &wireless, nil
}

const (
	EncryptionAlgTkip   = "tkip"
	EncryptionAlgCcmp   = "ccmp"
	EncryptionAlgWep40  = "wep40"
	EncryptionAlgWep104 = "wep104"
)

type Proto = string

const (
	ProtoAllowWPA2RSN Proto = "rsn"
	ProtoAllowWPA     Proto = "wpa"
)

type KeyManagement = string

const (
	KeyManagementNone             KeyManagement = "none"
	KeyManagementWPA2_3Personal   KeyManagement = "wpa-psk"             // WPA2 + WPA3 personal
	KeyManagementWPA3Personal     KeyManagement = "sae"                 // WPA3 personal only
	KeyManagementWPA2_3Enterprise KeyManagement = "wpa-eap"             // WPA2 + WPA3 enterprise
	KeyManagementWPA3Enterprise   KeyManagement = "wpa-eap-suite-b-192" // WPA3 enterprise only
)

type WirelessMode string

const (
	WirelessModeAccessPoint    WirelessMode = "ap"
	WirelessModeInfrastructure WirelessMode = "infrastructure"
	WirelessModeMesh           WirelessMode = "mesh"
	WirelessModeAdhoc          WirelessMode = "adhoc"
)

func (c *WirelessConnection) SetMode(mode WirelessMode) error {
	return c.setOption(OptionKeyWirelessMode, string(mode))
}

type WirelessBand = string

const (
	WirelessBand2GHz WirelessBand = "bg"
	WirelessBand5GHz WirelessBand = "a"
)

func (c *WirelessConnection) SetBand(band WirelessBand) error {
	return c.setOption(OptionKeyWirelessBand, string(band))
}
func (c *WirelessConnection) GetBand() WirelessBand {
	return WirelessBand(c.getOption(OptionKeyWirelessBand))
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
