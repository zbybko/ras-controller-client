package utils

import (
	"fmt"
	"os"
)

const RootUserID = 0

// Checks user is root, if not return error
func CheckRoot() error {
	if os.Geteuid() != RootUserID {
		return fmt.Errorf("current user is not root")
	}
	return nil
}
