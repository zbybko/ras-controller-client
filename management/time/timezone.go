package time

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/log"
)

func GetTimeZone() (string, error) {
	info, err := getInfo()
	if err != nil {
		return "", err
	}

	return info.GetTimeZone()
}

func SetTimeZone(timezone string) error {
	if err := checkRoot(); err != nil {
		log.Warn("Operation is not allowed for non-root users")
		return err
	}
	command := exec.Command("timedatectl", "set-timezone", inQuotes(timezone))
	err := command.Run()
	return err
}

func inQuotes(value string) string {
	return fmt.Sprintf("'%s'", value)
}
