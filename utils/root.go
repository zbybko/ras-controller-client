package utils

import (
	"fmt"
	"os"
	"os/exec"
	"ras/config"
)

const RootUserID = 0

// Checks user is root, if not return error
func CheckRoot() error {
	if os.Geteuid() != RootUserID {
		return fmt.Errorf("current user is not root")
	}
	return nil
}

func wrapCommand(command string, args ...string) string {
	return exec.Command(command, args...).String()
}

func Execute(command string, args ...string) ([]byte, error) {
	logger := config.GetLogger("CLI execution")

	wrapped := wrapCommand(command, args...)
	logger.Infof("Command to execute `%s`", wrapped)

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
