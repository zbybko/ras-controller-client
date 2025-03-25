package journals

import (
	"fmt"
	"ras/utils"

	"github.com/charmbracelet/log"
)

const JournalctlExecutable = "journalctl"
const KernelMessagesFlag = "--dmesg" // "-k"
const NoPagerFlag = "--no-pager"

var DefaultLines = 100

func linesFlag(lines int) string {
	return fmt.Sprintf("--lines=%d", lines)
}

func Core() (string, error) {
	output, err := utils.Execute(JournalctlExecutable, KernelMessagesFlag, linesFlag(DefaultLines), NoPagerFlag)
	return string(output), err
}

func System() (string, error) {
	output, err := utils.Execute(JournalctlExecutable, linesFlag(DefaultLines), NoPagerFlag)
	return string(output), err
}
func Connections() (string, error) {
	output, err := utils.Execute("ss", "-tuln")
	return string(output), err
}

// Returns port forwarding logs. Only if user is root
func PortForwarding() (string, error) {
	if err := utils.CheckRoot(); err != nil {
		log.Warnf("Can't get port forwarding logs: %s", err)
		return "", err
	}
	output, err := utils.Execute("iptables", "-t", "nat", "-L", "PREROUTING", "-n", "-v")
	return string(output), err
}
