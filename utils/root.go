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

func Execute(command string, args ...string) (string, error) {
	logger := config.GetLogger("CLI execution")

	cmd := exec.Command(command, args...)
	logger.Infof("Command to execute `%s`", cmd.String())
	output, err := cmd.Output()
	if err != nil {
		logger.Errorf("Failed execute command: %s", err)
	}
	return string(output), err
}
