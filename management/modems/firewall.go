package modems

import (
	"os/exec"
)

func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func Enable() error {
	_, err := executeCommand("systemctl", "start", "iptables")
	return err
}

func Disable() error {
	_, err := executeCommand("systemctl", "stop", "iptables")
	return err
}
