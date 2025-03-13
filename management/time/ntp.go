package time

import (
	"ras/management/time/chrony"
	"slices"

	"github.com/charmbracelet/log"
)

type TimeSyncManager interface {
	GetNtpServers() ([]string, error)
	AddNtpServer(string) error
	RemoveNtpServer(string) error
}

func GetNtpServers() ([]chrony.NtpServer, error) {

	config, err := chrony.ParseConfigFile(chrony.ChronyConfigFile)
	if err != nil {
		log.Warnf("Error while parsing server from chrony config: %s", err)
		return nil, err
	}

	return config.Servers, nil
}

func AddNtpServer(server *chrony.NtpServer) error {
	config, err := chrony.ParseConfigFile(chrony.ChronyConfigFile)
	if err != nil {
		log.Errorf("Error while parsing server from chrony config: %s", err)
		return err
	}
	config.Servers = append(config.Servers, *server)

	log.Debug("Saving chrony config")
	err = config.Save(chrony.ChronyConfigFile)

	if err != nil {
		log.Errorf("Error while writing chrony config: %s", err)
		return err
	}
	return chrony.Restart()
}
func RemoveNtpServer(server *chrony.NtpServer) error {
	config, err := chrony.ParseConfigFile(chrony.ChronyConfigFile)
	if err != nil {
		log.Errorf("Error while parsing server from chrony config: %s", err)
		return err
	}
	for i, srv := range config.Servers {
		if srv.Address() != server.Address() {
			continue
		}

		config.Servers = slices.Delete(config.Servers, i, i+1)
		break
	}

	log.Debug("Saving chrony config")
	err = config.Save(chrony.ChronyConfigFile)

	if err != nil {
		log.Errorf("Error while writing chrony config: %s", err)
		return err
	}
	return chrony.Restart()
}

func IsNtpActive() (bool, error) {
	info, err := getInfo()
	if err != nil {
		return false, err
	}
	return info.NTP(), nil
}
