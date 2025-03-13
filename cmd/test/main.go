package main

import (
	"ras/config"
	"ras/management/modems"
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

	// if err := time.AddNtpServer(chrony.NewNtpServer("ntp.katy248.ru")); err != nil {
	// 	log.Fatalf("Failed add NTP server: %s", err)
	// }

	// if err := time.RemoveNtpServer(chrony.NewNtpServer("ntp.katy248.ru")); err != nil {
	// 	log.Fatalf("Failed add NTP server: %s", err)
	// }

	// pass := storage.GetPassword()
	// log.Infof("Password hash: %s", pass)

	modemList, _ := modems.GetList()
	log.Infof("Modems: %v", modemList)
	for _, m := range modemList {
		mInfo, _ := modems.GetInfo(m)
		log.Infof("Modem: %v", mInfo)
	}
	log.Info("Done")
}
