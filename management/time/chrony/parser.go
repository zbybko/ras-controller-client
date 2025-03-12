package chrony

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var logger *log.Logger

func init() {
	logger = log.WithPrefix("Chrony config parser")
	if viper.GetBool("debug") {
		logger.SetLevel(log.DebugLevel)
	}
}

func ParseConfigFile(filename string) (*ChronyConfig, error) {

	file, err := os.Open(filename)
	if err != nil {
		logger.Errorf("Cannot open chrony config file: %s", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	config := ChronyConfig{}

	lineIndex := -1
	for scanner.Scan() {
		lineIndex++
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			logger.Debugf("Line %d: Empty", lineIndex)
			continue
		}
		if line[0] == '#' {
			logger.Debugf("Line %d: Empty", lineIndex)
			continue
		}
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case serverToken:
			server, err := parseServer(tokens)
			if err != nil {
				return &config, fmt.Errorf("failded to parse file, line %d: %s", lineIndex, err)
			}
			config.Servers = append(config.Servers, *server)
			logger.Debugf("Line %d: Server '%s'", lineIndex, server.Address())
		default:
			logger.Debugf("Line %d: Other parameter '%s'", lineIndex, tokens[0])
			param, err := parseTokens(tokens)
			if err != nil {
				return &config, fmt.Errorf("failded to parse file, line %d: %s", lineIndex, err)
			}
			config.Parameters = append(config.Parameters, *param)
		}
	}

	return &config, nil

}

func parseServer(tokens []string) (*NtpServer, error) {

	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid server specification, too little tokens")
	}
	if tokens[0] != serverToken {
		return nil, fmt.Errorf("not a server specification, first token must be a '%s' but there is '%s'", serverToken, tokens[0])
	}

	server := NtpServer{
		ChronyParameter{Name: serverToken, Value: tokens[1]},
	}
	if len(tokens) > 2 {
		server.Options = tokens[2:]
	}
	return &server, nil
}
func parseTokens(tokens []string) (*ChronyParameter, error) {
	param := ChronyParameter{}
	if len(tokens) >= 1 {
		param.Name = tokens[0]
	}
	if len(tokens) >= 2 {
		param.Value = tokens[1]
	}
	if len(tokens) > 2 {
		param.Options = tokens[:2]
	}

	return &param, nil

}

func (c *ChronyConfig) Save(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		log.Errorf("Error while writing chrony config: %s", err)
		return err
	}
	defer file.Close()
	return c.Dump(file)
}
