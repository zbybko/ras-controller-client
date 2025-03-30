package power

import (
	"ras/utils"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func Reboot() {
	if viper.GetBool("mock") {
		log.Info("Mock rebooting")
		log.Info("Mock rebooted")
		return
	}
	_, err := utils.Execute("reboot")
	if err != nil {
		log.Errorf("Failed reboot: %s", err)
	}
}
