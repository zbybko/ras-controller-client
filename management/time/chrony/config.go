package chrony

import (
	"fmt"
	"io"
	"os/exec"
	"ras/utils"
)

const (
	serverToken      = "server"
	poolToken        = "pool"
	ChronyConfigFile = "/etc/chrony.conf"
)

type TimeSyncServer interface {
	Address() string
	ToString() string
	Options() []string
}

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

type ChronyConfig struct {
	Servers    []TimeSyncServer
	Pools      []TimeSyncServer
	Parameters []ChronyParameter
}
type NtpServer struct {
	ChronyParameter
}
type NtpPool struct {
	ChronyParameter
}

func (s *NtpServer) Address() string {
	return s.Value
}
func (s *NtpServer) Options() []string {
	return s.ChronyParameter.Options
}
func (s *NtpPool) Address() string {
	return s.Value
}
func (s *NtpPool) Options() []string {
	return s.ChronyParameter.Options
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
