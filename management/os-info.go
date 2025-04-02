package management

import (
	"errors"
	"ras/management/nmcli"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

type OSInfo struct {
	Memory       *memory.Stats
	CpuStats     *cpu.Stats
	NetworkStats []NetworkStats
	DiskStats    []disk.Stats
	LoadAverage  *loadavg.Stats
}

func GetOSInfo() (OSInfo, error) {

	mem, memErr := memory.Get()
	cpuS, cpuErr := cpu.Get()
	net, netErr := network.Get()
	disk, diskErr := disk.Get()
	load, loadErr := loadavg.Get()
	if err := errors.Join(memErr, cpuErr, netErr, diskErr, loadErr); err != nil {
		return OSInfo{}, err
	}
	return OSInfo{
		CpuStats:     cpuS,
		NetworkStats: toNetworkStatsArr(net),
		Memory:       mem,
		DiskStats:    disk,
		LoadAverage:  load}, nil
}

type NetworkStats struct {
	network.Stats
	MAC string `json:"MAC"`
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
