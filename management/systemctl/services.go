package systemctl

import (
	"os/exec"
	"ras/utils"
	"strings"

	"github.com/charmbracelet/log"
)

const SystemctlExecutable = "systemctl"

const ExitCodeInactive = 3

const (
	StatusActive = "active"
)

func ServiceExists(name string) bool {
	_, err := utils.Execute(SystemctlExecutable, "status", name)
	return err != nil
}
func Enable(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "enable", "--now", name)
	return err
}
func Disable(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "disable", "--now", name)
	return err
}

func IsActive(name string) bool {
	output, err := utils.Execute(SystemctlExecutable, "is-active", name)

	// Pass if error is 'exit code 3'
	// Exit code 3 is for inactive state of service
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() != ExitCodeInactive {
				log.Errorf("Failed get status of '%s' service, exit code isn't valid: %s", name, err)
			}
		} else {
			log.Errorf("Failed get status of '%s' service: %s", name, err)
			return false

		}
	}
	strOutput := strings.TrimSpace(string(output))
	return strOutput == StatusActive
}
func Restart(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "restart", name)
	return err
}
