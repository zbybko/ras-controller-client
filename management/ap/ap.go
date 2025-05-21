package ap

import (
	"ras/management/nmcli"

	"github.com/charmbracelet/log"
)

const ConnectionName = "HostspotZarinit"

type Info struct {
	Channel  int    `json:"channel"`
	SSID     string `json:"ssid"`
	Password string `json:"password"`
	Hidden   bool   `json:"hidden"`
}

func Status() (*Info, error) {
	conn, err := getConnection()

	if err != nil {
		log.Errorf("Failed get connection: %s", err)
		return nil, err
	}
	return getConnectionInfo(conn)
}

func getConnectionInfo(conn *nmcli.WirelessConnection) (*Info, error) {
	info := Info{
		Channel:  conn.GetChanel(),
		SSID:     conn.GetSSID(),
		Password: conn.GetPassword(),
		Hidden:   conn.IsHidden(),
	}

	return &info, nil
}

func Enable() error {
	conn, err := getConnection()

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

func Disable() error {
	conn, err := getConnection()

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
func SetSSID(ssid string) error {
	conn, err := getConnection()

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
func SetPassword(password string) error {
	conn, err := getConnection()

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
func SetHidden(hidden bool) error {
	conn, err := getConnection()

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
func SetChannel(channel int) error {
	conn, err := getConnection()

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

// type ConnectedDevice struct {
// 	MAC       string `json:"mac"`
// 	Interface string `json:"interface"`
// 	Signal    string `json:"signal"`
// }

// func GetConnectedDevices() {

// 	output, err := utils.Execute("iw", "dev")
// }
