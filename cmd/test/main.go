package main

import (
	"ras/config"
	"ras/management/time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func init() {

}

func main() {
	config.LoadConfigFile()
	config.SetupLogger()
	log.Info(viper.GetString("config"))
	t, _ := time.GetTimeZone()
	log.Infof("Timezone: %s", t)
	ntpActive, _ := time.IsNtpActive()
	log.Infof("Ntp active: %v", ntpActive)
	time.SetTimeZone("Europe/Moscow")
	ntpServers, _ := time.GetNtpServers()
	log.Infof("Ntp server: %s", ntpServers)
}
