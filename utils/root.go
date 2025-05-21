package utils

import (
	"fmt"
	"os"
	"os/exec"
	"ras/config"

	"github.com/charmbracelet/log"
)

const RootUserID = 0

var ErrNotRootUser = fmt.Errorf("current user is not root")

// Checks user is root, if not return error
func CheckRoot() error {
	if os.Geteuid() != RootUserID {
		return ErrNotRootUser
	}
	return nil
}

func wrapCommand(command string, args ...string) string {
	return exec.Command(command, args...).String()
}

func execute(wrapInQuotes bool, command string, args ...string) ([]byte, error) {
	logger := config.GetLogger("CLI execution")

	wrapped := wrapCommand(command, args...)
	logger.Infof("Command to execute `%s`", wrapped)

	if wrapInQuotes {
		log.Debug("Wrapping command in quotes")
		wrapped = fmt.Sprintf("\"%s\"", wrapped)
	}

	cmd := exec.Command("bash", "-c", wrapped)
	logger.Debugf("Real command to execute `%s`", cmd.String())

	output, err := cmd.Output()
	if err != nil {
		logger.Warnf("Error while executing command `%s` execute command: %s", cmd.String(), err)
		logger.Debugf("Output: %s", string(output))
		return nil, err
	}
	return output, err

}

func Execute(command string, args ...string) ([]byte, error) {
	return execute(false, command, args...)
}
