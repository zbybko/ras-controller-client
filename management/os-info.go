package management

import (
	"errors"
	"ras/management/nmcli"
	"ras/utils"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

type NetworkStats struct {
	network.Stats
	MAC string `json:"MAC"`
}

// TODO: move string fields to integers
type DiskStats struct {
	Name      string `json:"name"`
	Size      string `json:"size"`
	Used      string `json:"used"`
	Available string `json:"available"`
	MountPint string `json:"mountPoint"`
}
type OSInfo struct {
	Memory       *memory.Stats
	CpuStats     *cpu.Stats
	NetworkStats []NetworkStats
	DiskStats    []DiskStats
	LoadAverage  *loadavg.Stats
}

func GetOSInfo() (OSInfo, error) {

	mem, memErr := memory.Get()
	cpuS, cpuErr := cpu.Get()
	net, netErr := network.Get()
	load, loadErr := loadavg.Get()
	if err := errors.Join(memErr, cpuErr, netErr, loadErr); err != nil {
		return OSInfo{}, err
	}
	return OSInfo{
		CpuStats:     cpuS,
		NetworkStats: toNetworkStatsArr(net),
		Memory:       mem,
		DiskStats:    getDiskStats(),
		LoadAverage:  load}, nil
}

func toNetworkStatsArr(stats []network.Stats) []NetworkStats {
	stats2 := make([]NetworkStats, len(stats))
	for _, s := range stats {
		stats2 = append(stats2, newNetworkStats(s))
	}
	return stats2

}

func newNetworkStats(s network.Stats) NetworkStats {
	mac, _ := nmcli.GetHardwareAddress(s.Name)
	return NetworkStats{
		Stats: s,
		MAC:   mac,
	}
}

func getDiskStats() []DiskStats {
	val, err := utils.Execute("df", "--exclude-type=tmpfs", "--exclude-type=squashfs", "--exclude-type=devtmpfs", "--output=source,size,used,avail,target")
	if err != nil {
		log.Errorf("Failed to get disk stats: %v", err)
		return []DiskStats{}
	}

	lines := strings.Split(
		strings.TrimSpace(string(val)),
		"\n")[1:]

	result := []DiskStats{}
	for _, line := range lines {
		var stats DiskStats
		fields := strings.Fields(line)
		stats.Name = fields[0]
		stats.Size = fields[1]
		stats.Used = fields[2]
		stats.Available = fields[3]
		stats.MountPint = fields[4]

		result = append(result, stats)
	}

	return result
}
