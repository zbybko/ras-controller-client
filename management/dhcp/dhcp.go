package dhcp

import (
	"bufio"
	"errors"
	"os"
	"ras/management/systemctl"
	"ras/utils"
	"strings"
	"time"
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

const DhcpService = "dhcpd.service"
const LeaseFile = "/var/lib/dhcpd/dhcpd.leases"

func Status() *DhcpStatus {
	return &DhcpStatus{
		Enabled: systemctl.IsActive(DhcpService),
	}
}

func Enable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	return errors.Join(
		systemctl.Enable(DhcpService),
		utils.ExecuteErr("firewall-cmd", "--permanent", "--add-service=dhcp"),
		utils.ExecuteErr("firewall-cmd", "--reload"),
	)
}

func Disable() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	return systemctl.Disable(DhcpService)
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
