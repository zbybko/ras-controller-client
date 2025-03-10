package time

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	l "github.com/charmbracelet/log"
)

var logger = l.New(os.Stdout)

func init() {
	logger.SetPrefix("NTP management")
}

const ChronyConfigFile = "/etc/chrony.conf"

func GetNtpServers() ([]NtpServer, error) {
	logger.SetPrefix("NTP management")

	servers := []NtpServer{}
	file, err := os.Open(ChronyConfigFile)
	if err != nil {

		logger.Errorf("Cannot open chrony config file: %s", err)
		return servers, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	parserLogger := l.New(os.Stdout)
	parserLogger.SetPrefix("Chrony config parser")

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			parserLogger.Debug("Empty line")
			continue
		}
		if line[0] == '#' {
			parserLogger.Debug("Comment line")
			continue
		}
		server, err := parseServerLine(line)
		if err != nil {
			parserLogger.Warnf("Error while parsing server from chrony config: %s", err)
			continue
		}
		servers = append(servers, *server)
	}

	return servers, nil
}

type NtpServer struct {
	Address string
	Options []string
}

const (
	serverToken = "server"
)

func parseServerLine(line string) (*NtpServer, error) {
	if line == "" {
		return nil, fmt.Errorf("empty line")
	}

	tokens := strings.Split(line, " ")

	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid server specification, too little tokens")
	}
	if tokens[0] != serverToken {
		return nil, fmt.Errorf("not a server specification, first token must be a '%s' but there is '%s'", serverToken, tokens[0])
	}

	server := NtpServer{
		Address: tokens[1],
	}
	if len(tokens) > 2 {
		server.Options = tokens[2:]
	}
	return &server, nil
}
func IsNtpActive() (bool, error) {
	info, err := getInfo()
	if err != nil {
		return false, err
	}
	return info.NTP(), nil
}
