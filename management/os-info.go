package management

import (
	"errors"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

type MemoryInfo struct {
	Total uint64
	Used  uint64
}
type OSInfo struct {
	Memory       MemoryInfo
	CpuStats     *cpu.Stats
	NetworkStats []network.Stats
	DiskStats    []disk.Stats
	LoadAverage  loadavg.Stats
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
		NetworkStats: net,
		Memory: MemoryInfo{
			Total: mem.Total,
			Used:  mem.Used,
		},
		DiskStats:   disk,
		LoadAverage: *load}, nil

}
