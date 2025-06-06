package wifi

import (
	"ras/management/nmcli"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func getConnectionName(band nmcli.WirelessBand) string {
	return "ZarinitAccessPoint" + "-" + string(band)

}

func getConnection(band nmcli.WirelessBand) (*nmcli.WirelessConnection, error) {
	exists, conn := connectionExists(band)
	if exists {
		return conn, nil
	}

	return createConnection(band)
}

func connectionExists(band nmcli.WirelessBand) (bool, *nmcli.WirelessConnection) {
	connections, err := nmcli.GetConnections()
	if err != nil {
		log.Errorf("Failed get connection list: %s", err)
		return false, nil
	}
	for _, conn := range connections {
		if conn.Name == getConnectionName(band) {
			return true, &nmcli.WirelessConnection{Connection: &conn}
		}
	}

	log.Warnf("No connection for band %s", band)
	return false, nil
}
func createConnection(band nmcli.WirelessBand) (*nmcli.WirelessConnection, error) {
	interfaceName := viper.GetString("wifi.default_interface")

	conn, err := nmcli.CreateWirelessConnection(interfaceName, getConnectionName(band))
	conn.SetMode(nmcli.WirelessModeAccessPoint)
	conn.SetBand(band)
	conn.SetIP4Method(nmcli.ConnectionIP4MethodShared)

	if err != nil {
		log.Errorf("Failed create new connection: %s", err)
		return nil, err
	}
	return conn, nil
}
