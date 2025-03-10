package time

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
)

type TimedateInfoDictionary map[string]string

const RootUserID = 0

// Checks user is root, if not return error
func checkRoot() error {
	if os.Geteuid() != RootUserID {
		return fmt.Errorf("current user is not root")
	}
	return nil
}

func getInfo() (TimedateInfoDictionary, error) {
	dict := TimedateInfoDictionary{}
	command := exec.Command("timedatectl", "show")
	output, err := command.Output()
	if err != nil {
		return dict, fmt.Errorf("failed get output from timedatectl: %s", err)
	}
	outputStr := string(output)
	for _, line := range strings.Split(outputStr, "\n") {
		keyValArr := strings.Split(line, "=")
		if len(keyValArr) < 2 {
			continue
		}

		dict[keyValArr[0]] = keyValArr[1]
	}

	return dict, nil
}

func (info TimedateInfoDictionary) GetTimeZone() (string, error) {
	var err error
	timezone, ok := info[timeZoneKey]
	if !ok {
		err = fmt.Errorf("failed to get timezone info")
	}
	return timezone, err
}

func (info TimedateInfoDictionary) NTP() bool {
	ntp, ok := info[ntpKey]
	if !ok {
		log.Errorf("failed to get timezone info")
		return false
	}
	return ntp == "yes"

}

const (
	timeZoneKey = "Timezone"
	ntpKey      = "NTP"
)
