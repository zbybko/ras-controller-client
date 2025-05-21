package wifi

import (
	"ras/management/nmcli"

	"github.com/charmbracelet/log"
)

func getConnection() (*nmcli.WirelessConnection, error) {
	exists, conn := connectionExists()
	if exists {
		return conn, nil
	}

	return createConnection()
}

func connectionExists() (bool, *nmcli.WirelessConnection) {
	connections, err := nmcli.GetConnections()
	if err != nil {
		log.Errorf("Failed get connection list: %s", err)
		return false, nil
	}
	for _, conn := range connections {
		if conn.Name == ConnectionName {
			return true, &nmcli.WirelessConnection{Connection: &conn}
		}
	}

	log.Warnf("No connection")
	return false, nil
}
func createConnection() (*nmcli.WirelessConnection, error) {
	const interfaceName = "wlan0" // TODO: fix default interface/device

	log.Warnf("Creating new connection with default interface %s, this should be fixed ASAP", interfaceName)
	log.Warnf("Immediately call function author Katy248 to fix this")

	conn, err := nmcli.CreateWirelessConnection(interfaceName, ConnectionName)
	conn.SetMode(nmcli.WirelessModeAccessPoint)
	conn.SetBand(nmcli.WirelessBandDefault)
	conn.SetIP4Method(nmcli.ConnectionIP4MethodShared)

	if err != nil {
		log.Errorf("Failed create new connection: %s", err)
		return nil, err
	}
	return conn, nil
}
