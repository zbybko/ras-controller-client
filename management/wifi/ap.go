package wifi

import (
	"ras/management/nmcli"

	"github.com/charmbracelet/log"
)

type Info struct {
	Channel  int                `json:"channel"`
	SSID     string             `json:"ssid"`
	Password string             `json:"password"`
	Hidden   bool               `json:"hidden"`
	Active   bool               `json:"active"`
	Band     nmcli.WirelessBand `json:"band"`
}

func Status(band nmcli.WirelessBand) (*Info, error) {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return nil, err
	}
	return getConnectionInfo(conn)
}

func getConnectionInfo(conn *nmcli.WirelessConnection) (*Info, error) {
	info := Info{
		Active:   conn.IsActive(),
		Channel:  conn.GetChanel(),
		SSID:     conn.GetSSID(),
		Password: conn.GetPassword(),
		Hidden:   conn.IsHidden(),
		Band:     conn.GetBand(),
	}

	return &info, nil
}

func Enable(band nmcli.WirelessBand) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.Up()
	if err != nil {
		log.Errorf("Failed enable connection: %s", err)
		return err
	}
	return nil
}

func Disable(band nmcli.WirelessBand) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.Down()
	if err != nil {
		log.Errorf("Failed disable connection: %s", err)
		return err
	}
	return nil
}
func SetSSID(band nmcli.WirelessBand, ssid string) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.SetSSID(ssid)
	if err != nil {
		log.Errorf("Failed set ssid: %s", err)
		return err
	}
	return nil
}
func SetPassword(band nmcli.WirelessBand, password string) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.SetPassword(password)
	if err != nil {
		log.Errorf("Failed set password: %s", err)
		return err
	}
	return nil
}
func SetHidden(band nmcli.WirelessBand, hidden bool) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.SetHidden(hidden)
	if err != nil {
		log.Errorf("Failed set hidden: %s", err)
		return err
	}
	return nil
}
func SetChannel(band nmcli.WirelessBand, channel int) error {
	conn, err := getConnection(band)

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return err
	}

	err = conn.SetChannel(channel)
	if err != nil {
		log.Errorf("Failed set channel: %s", err)
		return err
	}
	return nil
}
