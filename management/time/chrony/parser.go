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
		if isCommentSymbol(line[0]) {
			logger.Debugf("Line %d: Commented", lineIndex)
			continue
		}
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case serverToken:
			server, err := parseServer(tokens)
			if err != nil {
				return &config, fmt.Errorf("failed to parse file, line %d: %s", lineIndex, err)
			}
			config.Servers = append(config.Servers, server)
			logger.Debugf("Line %d: Server '%s'", lineIndex, server.Address())
		case poolToken:
			pool, err := parsePool(tokens)
			if err != nil {
				return &config, fmt.Errorf("failed to parse file, line %d: %s", lineIndex, err)
			}
			config.Pools = append(config.Pools, pool)
			logger.Debugf("Line %d: Pool '%s'", lineIndex, pool.Address())
		default:
			logger.Debugf("Line %d: Other parameter '%s'", lineIndex, tokens[0])
			param, err := parseChronyParameter(tokens, "", 0)
			if err != nil {
				return &config, fmt.Errorf("failed to parse file, line %d: %s", lineIndex, err)
			}
			config.Parameters = append(config.Parameters, *param)
		}
	}

	return &config, nil
}

func isCommentSymbol(sym byte) bool {
	return sym == '#' || sym == '!' || sym == '%' || sym == ';'
}
func parsePool(tokens []string) (*NtpPool, error) {
	param, err := parseChronyParameter(tokens, poolToken, 2)

	if err != nil {
		return nil, err
	}

	pool := NtpPool{
		ChronyParameter: *param,
	}
	return &pool, nil
}

// pass name = "" to skip check
// pass minLength = {<=0} to skip check
func parseChronyParameter(tokens []string, name string, minLength int) (*ChronyParameter, error) {
	if name != "" && tokens[0] != name {
		return nil, fmt.Errorf("invalid parameter specification, first token must be a '%s' but there is '%s'", name, tokens[0])
	}
	if minLength > 0 && len(tokens) < minLength {
		return nil, fmt.Errorf("invalid parameter specification, too little tokens")
	}
	param := ChronyParameter{}
	if len(tokens) >= 1 {
		param.Name = tokens[0]
	}
	if len(tokens) >= 2 {
		param.Value = tokens[1]
	}
	if len(tokens) > 2 {
		param.Options = tokens[2:]
	}

	return &param, nil
}
func parseServer(tokens []string) (*NtpServer, error) {
	param, err := parseChronyParameter(tokens, serverToken, 2)
	if err != nil {
		return nil, err
	}
	server := NtpServer{
		ChronyParameter: *param,
	}
	return &server, nil
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
