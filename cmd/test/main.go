package main

import (
	"ras/config"
	"ras/management/firewall"
	"ras/management/journals"
	"ras/management/modems"
	"ras/management/time"
	"ras/storage"

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

	modemList, _ := modems.List()
	log.Infof("Modems: %v", modemList)
	for _, m := range modemList {
		mInfo, _ := modems.Get(m)
		log.Infof("Modem: %v", mInfo)
	}
	log.Info("Done")
	// FirewallTest()
	// TestPasswordStorage()
	JournalsTest()
}
func JournalsTest() {
	core, err := journals.Core()
	if err != nil {
		log.Errorf("Failed get journals: %s", err)
		return
	}
	log.Infof("Journals: %v", core)
	sys, err := journals.System()
	if err != nil {
		log.Errorf("Failed get journals: %s", err)
		return
	}
	log.Infof("Journals: %v", sys)
}

func TestPasswordStorage() {
	pass := storage.GetPassword()
	log.Infof("Password hash: %s", pass)
}

func FirewallTest() {
	firewallInfo, err := firewall.Status()
	if err != nil {
		log.Errorf("Failed get firewall info: %s", err)
		return
	}
	if !firewallInfo.Active {
		log.Debug("Firewall inactive")
		firewall.Enable()
	} else {
		log.Debug("Firewall active")
		firewall.Disable()
	}

}
