package ap

import (
	"ras/management/nmcli"

	"github.com/charmbracelet/log"
)

const ConnectionName = "HostspotZarinit"

func connectionExists() (bool, *nmcli.Connection) {
	connections, err := nmcli.GetConnections()
	if err != nil {
		log.Errorf("Failed get connection list: %s", err)
		return false, nil
	}
	for _, conn := range connections {
		if conn.Name == ConnectionName {
			return true, &conn
		}
	}

	log.Warnf("No connection")
	return false, nil
}

// func GetConnection() (*nmcli.Connection, error) {
// 	exists, conn := connectionExists()
// 	if exists {
// 		return conn, nil
// 	}
// 	conn, err := nmcli.CreateConnection()

// 	if err != nil {
// 		log.Errorf("Failed create new connection: %s", err)
// 		return nil, err
// 	}
// 	return conn, nil
// }

// type ConnectedDevice struct {
// 	MAC       string `json:"mac"`
// 	Interface string `json:"interface"`
// 	Signal    string `json:"signal"`
// }

// func GetConnectedDevices() {

// 	output, err := utils.Execute("iw", "dev")
// }
