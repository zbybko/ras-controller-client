package ethernet

import (
	"net"
	"strings"
)

type EthernetPort struct {
	Name string `json:"name"`
	MAC  string `json:"mac"`
	IP   string `json:"ip,omitempty"`
	Up   bool   `json:"up"`
}

func GetEthernetPorts() ([]EthernetPort, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ports []EthernetPort

	for _, iface := range ifaces {
		// Отбираем только Ethernet-порты
		if !(strings.HasPrefix(iface.Name, "en") || strings.HasPrefix(iface.Name, "eth")) {
			continue
		}

		// Получаем IP-адрес (если есть)
		addrs, _ := iface.Addrs()
		var ip string
		for _, addr := range addrs {
			if ip == "" && strings.Contains(addr.String(), ".") {
				ip = strings.Split(addr.String(), "/")[0]
			}
		}

		ports = append(ports, EthernetPort{
			Name: iface.Name,
			MAC:  iface.HardwareAddr.String(),
			IP:   ip,
			Up:   iface.Flags&net.FlagUp != 0,
		})
	}

	return ports, nil
}
