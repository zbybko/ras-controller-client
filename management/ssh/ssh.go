package ssh

import (
	"ras/management/systemctl"
	"ras/utils"
	"strings"
)

type SshStatus struct {
	Enabled bool `json:"enabled"`
}

const SshService = "sshd"

func Status() SshStatus {
	sshEnabled := systemctl.IsActive(SshService)

	return SshStatus{
		Enabled: sshEnabled,
	}
}

func Enable() error {
	return systemctl.Enable(SshService)
}
func Disable() error {
	return systemctl.Disable(SshService)
}

func GetKeys() ([]string, error) {
	output, err := utils.Execute("ssh-keyscan", "localhost")
	if err != nil {
		return nil, err
	}
	strs := strings.Split(string(output), "\n")
	keys := make([]string, len(strs))
	for _, s := range strs {
		if s[0] != '#' {
			keys = append(keys, s)
		}
	}

	return keys, nil
}
