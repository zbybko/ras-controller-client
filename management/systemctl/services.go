package systemctl

import (
	"os/exec"
	"ras/config"
	"ras/utils"
	"strings"

	l "github.com/charmbracelet/log"
)

const SystemctlExecutable = "systemctl"

const ExitCodeInactive = 3

const (
	StatusActive = "active"
)

var log *l.Logger

func init() {
	log = config.GetLogger("Systemctl module")
}

func ServiceExists(name string) bool {
	_, err := utils.Execute(SystemctlExecutable, "list-unit-files", name)
	return err == nil
}
func Enable(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "enable", "--now", name)
	if err != nil {
		log.Errorf("Failed enable '%s' service: %s", name, err)
		printErrorDebugInfo(name)
	}
	return err
}
func Disable(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "disable", "--now", name)
	if err != nil {
		log.Errorf("Failed disable '%s' service: %s", name, err)
		printErrorDebugInfo(name)
	}
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
		printErrorDebugInfo(name)
	}
	strOutput := strings.TrimSpace(string(output))
	return strOutput == StatusActive
}
func Restart(name string) error {
	_, err := utils.Execute(SystemctlExecutable, "restart", name)
	if err != nil {
		log.Errorf("Failed restart '%s' service: %s", name, err)
		printErrorDebugInfo(name)
	}
	return err
}
func printErrorDebugInfo(serviceName string) {
	log.Debugf("See `journalctl -xeu %s`", serviceName)
}
