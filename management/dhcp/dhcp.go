package dhcp

import (
	"bufio"
	"errors"
	"os"
	"ras/management/systemctl"
	"ras/utils"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type DhcpStatus struct {
	Enabled bool `json:"enabled"`
}

type Lease struct {
	IP       string    `json:"ip"`
	MAC      string    `json:"mac"`
	Hostname string    `json:"hostname"`
	Expires  time.Time `json:"expires"`
}

type StaticLease struct {
	Hostname string `json:"hostname"`
	MAC      string `json:"mac"`
	IP       string `json:"ip"`
}

type Range struct {
	Subnet            string `json:"subnet"`
	Netmask           string `json:"netmask"`
	StartIP           string `json:"start_ip"`
	EndIP             string `json:"end_ip"`
	OptionsRouters    string `json:"options_routers"`
	OptionsBroadcasts string `json:"options_broadcasts"`
}

const DhcpService = "dhcpd.service"   // Default service for RED OS 8
const Dhcp4Service = "dhcpd4.service" // Default service for Arch Linux
const LeaseFile = "/var/lib/dhcpd/dhcpd.leases"
const DhcpConfig = "/etc/dhcp/dhcpd.conf"

var ErrNoDhcpService = errors.New("no dhcp service found")

func getCurrentService() (string, error) {
	if systemctl.ServiceExists(DhcpService) {
		log.Debug("DHCP service found", "service", DhcpService)
		return DhcpService, nil
	}
	if systemctl.ServiceExists(Dhcp4Service) {
		log.Debug("DHCP service found", "service", Dhcp4Service)
		return Dhcp4Service, nil
	}
	log.Warn("No DHCP service found")
	return "", ErrNoDhcpService
}

func Status() *DhcpStatus {
	service, err := getCurrentService()
	if err != nil {
		log.Errorf("Failed get dhcp service: %s", err)
		return nil
	}
	return &DhcpStatus{
		Enabled: systemctl.IsActive(service),
	}
}

func Enable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	service, err := getCurrentService()
	if err != nil {
		return err
	}
	return errors.Join(
		systemctl.Enable(service),
		utils.ExecuteErr("firewall-cmd", "--permanent", "--add-service=dhcp"),
		utils.ExecuteErr("firewall-cmd", "--reload"),
	)
}

func Disable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	service, err := getCurrentService()
	if err != nil {
		return err
	}
	return systemctl.Disable(service)
}

func CanManage() bool {
	if _, err := getCurrentService(); errors.Is(err, ErrNoDhcpService) {
		return false
	}
	return true
}

func GetLeases() ([]Lease, error) {
	file, err := os.Open(LeaseFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var leases []Lease
	scanner := bufio.NewScanner(file)

	var lease Lease
	inLease := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "lease ") {
			lease = Lease{}
			lease.IP = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "lease"), "{"))
			inLease = true
			continue
		}
		if line == "}" && inLease {
			if lease.IP != "" && lease.MAC != "" {
				leases = append(leases, lease)
			}
			inLease = false
			continue
		}
		if !inLease {
			continue
		}

		if strings.HasPrefix(line, "ends ") {
			lease.Expires = parseLeaseTime(line)
		} else if strings.HasPrefix(line, "hardware ethernet ") {
			lease.MAC = strings.TrimSuffix(strings.TrimPrefix(line, "hardware ethernet "), ";")
		} else if strings.HasPrefix(line, "client-hostname ") {
			lease.Hostname = strings.Trim(strings.TrimSuffix(strings.TrimPrefix(line, "client-hostname "), ";"), `"`)
		}
	}

	return leases, scanner.Err()
}

func parseLeaseTime(line string) time.Time {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return time.Time{}
	}
	layout := "2006/01/02 15:04:05"
	t, err := time.Parse(layout, parts[2]+" "+strings.TrimSuffix(parts[3], ";"))
	if err != nil {
		return time.Time{}
	}
	return t
}

func SetDhcpRange(subnet, netmask, startIP, endIP, routerIP, broadcastIP string) error {
	data, err := os.ReadFile(DhcpConfig)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	var subnetAndNetmaskUpdated, rangeUpdated, routerUpdated, broadcastUpdated bool

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(trimmed, "subnet "):
			lines[i] = "subnet " + subnet + " netmask " + netmask + "{"
			subnetAndNetmaskUpdated = true

		case strings.HasPrefix(trimmed, "range "):
			indent := line[:strings.Index(line, "r")]
			lines[i] = indent + "range " + startIP + " " + endIP + ";"
			rangeUpdated = true

		case strings.HasPrefix(trimmed, "option routers "):
			indent := line[:strings.Index(line, "o")]
			lines[i] = indent + "option routers " + routerIP + ";"
			routerUpdated = true

		case strings.HasPrefix(trimmed, "option broadcast-address "):
			indent := line[:strings.Index(line, "o")]
			lines[i] = indent + "option broadcast-address " + broadcastIP + ";"
			broadcastUpdated = true
		}
	}

	if !(rangeUpdated && routerUpdated && broadcastUpdated && subnetAndNetmaskUpdated) {
		return errors.New("one or more DHCP options were not found in the config")
	}

	newContent := strings.Join(lines, "\n")
	return os.WriteFile(DhcpConfig, []byte(newContent), 0644)
}

func GetDhcpRange() (*Range, error) {
	data, err := os.ReadFile(DhcpConfig)
	if err != nil {
		return nil, err
	}

	var startIP, endIP, optionsRouters, optionsBroadcasts, subnet, netmask string
	content := string(data)

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trimmed, "subnet "):
			parts := strings.Fields(trimmed)
			if len(parts) >= 4 {
				subnet = parts[1]
				netmask = parts[3]
			}
		case strings.HasPrefix(trimmed, "range "):
			parts := strings.Fields(trimmed)
			if len(parts) >= 3 {
				startIP = parts[1]
				endIP = strings.TrimSuffix(parts[2], ";")
			}
		case strings.HasPrefix(trimmed, "option routers "):
			parts := strings.Fields(trimmed)
			if len(parts) >= 3 {
				optionsRouters = strings.TrimSuffix(parts[2], ";")
			}
		case strings.HasPrefix(trimmed, "option broadcast-address "):
			parts := strings.Fields(trimmed)
			if len(parts) >= 3 {
				optionsBroadcasts = strings.TrimSuffix(parts[2], ";")
			}
		}
	}

	if startIP == "" && endIP == "" {
		return nil, errors.New("DHCP range not found in the configuration file")
	}

	return &Range{
		Subnet:            subnet,
		Netmask:           netmask,
		StartIP:           startIP,
		EndIP:             endIP,
		OptionsRouters:    optionsRouters,
		OptionsBroadcasts: optionsBroadcasts,
	}, nil

}

func AddStaticLease(mac, ip, hostname string) error {
	data, err := os.ReadFile(DhcpConfig)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	var leaseAdded bool

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "host ") {
			if strings.Contains(line, mac) {
				return nil
			}
		}
	}

	staticLease := "\nhost " + hostname + " {\n  hardware ethernet " + mac + ";\n  fixed-address " + ip + ";\n}"
	lines = append(lines, staticLease)
	leaseAdded = true

	if leaseAdded {
		newContent := strings.Join(lines, "\n")
		err := os.WriteFile(DhcpConfig, []byte(newContent), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveStaticLease(mac string) error {
	data, err := os.ReadFile(DhcpConfig)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	var leaseRemoved bool

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "host ") {
			blockStart := i
			blockEnd := -1
			foundMac := false

			for j := i + 1; j < len(lines); j++ {
				trimmed := strings.TrimSpace(lines[j])

				if strings.HasPrefix(trimmed, "hardware ethernet") {
					macInFile := strings.TrimSuffix(strings.TrimPrefix(trimmed, "hardware ethernet "), ";")
					if strings.EqualFold(macInFile, mac) {
						foundMac = true
					}
				}

				if trimmed == "}" {
					blockEnd = j
					break
				}
			}

			if foundMac && blockEnd != -1 {
				lines = append(lines[:blockStart], lines[blockEnd+1:]...)
				leaseRemoved = true
				break
			}
		}
	}

	if !leaseRemoved {
		return errors.New("static lease with the given MAC address not found")
	}

	newContent := strings.Join(lines, "\n")
	return os.WriteFile(DhcpConfig, []byte(newContent), 0644)
}

func GetStaticLeases() ([]StaticLease, error) {
	data, err := os.ReadFile(DhcpConfig)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var leases []StaticLease

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "host ") {
			var lease StaticLease
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				lease.Hostname = parts[1]
			}

			lease.Hostname = strings.TrimSuffix(lease.Hostname, "{")

			for j := i + 1; j < len(lines); j++ {
				trimmed := strings.TrimSpace(lines[j])
				if strings.HasPrefix(trimmed, "hardware ethernet ") {
					lease.MAC = strings.TrimSuffix(strings.TrimPrefix(trimmed, "hardware ethernet "), ";")
				} else if strings.HasPrefix(trimmed, "fixed-address ") {
					lease.IP = strings.TrimSuffix(strings.TrimPrefix(trimmed, "fixed-address "), ";")
				} else if trimmed == "}" {
					break
				}
			}

			if lease.MAC != "" && lease.IP != "" {
				leases = append(leases, lease)
			}
		}
	}

	return leases, nil
}

func RestartDhcp() error {
	service, err := getCurrentService()
	if err != nil {
		return err
	}
	return systemctl.Restart(service)
}
