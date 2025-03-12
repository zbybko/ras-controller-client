package chrony

import (
	"fmt"
	"io"
	"os/exec"
	"ras/utils"
)

const (
	serverToken      = "server"
	ChronyConfigFile = "/etc/chrony.conf"
)

func Restart() error {
	if err := utils.CheckRoot(); err != nil {
		return err
	}
	cmd := exec.Command("systemctl", "restart", "chronyd.service")
	return cmd.Run()
}

type ChronyParameter struct {
	Name    string
	Value   string
	Options []string
}

// func (c1 *ChronyParameter) Compare(c2 *ChronyParameter) bool {
// 	return c1 == c2
// }

type ChronyConfig struct {
	Servers    []NtpServer
	Parameters []ChronyParameter
}
type NtpServer struct {
	ChronyParameter
}

func (s *NtpServer) Address() string {
	return s.Value
}
func NewNtpServer(address string) *NtpServer {
	param := ChronyParameter{
		Name:  serverToken,
		Value: address,
		// Options: []string{"iburst"}
	}
	server := &NtpServer{param}
	return server
}

func (p *ChronyParameter) ToString() string {
	str := fmt.Sprintf("%s %s", p.Name, p.Value)
	for _, opt := range p.Options {
		str += " " + opt
	}
	return str
}

func (c *ChronyConfig) Dump(file io.Writer) error {
	for _, server := range c.Servers {
		_, err := fmt.Fprintf(file, "%s\n", server.ToString())
		if err != nil {
			return fmt.Errorf("failed to write config line: %s", err)
		}
	}
	for _, p := range c.Parameters {
		_, err := fmt.Fprintf(file, "%s\n", p.ToString())
		if err != nil {
			return fmt.Errorf("failed to write config line: %s", err)
		}
	}
	return nil
}
