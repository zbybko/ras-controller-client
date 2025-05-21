package time

import (
	"ras/utils"

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
	if err := utils.CheckRoot(); err != nil {
		log.Warn("Operation is not allowed for non-root users")
		return err
	}
	_, err := utils.Execute("timedatectl", "set-timezone", timezone)
	return err
}
