package diagnostics

import (
	"fmt"
	"ras/utils"

	"github.com/charmbracelet/log"
)

var ErrEmptyAddress = fmt.Errorf("empty address string")

func Nslookup(addr string) (string, error) {
	if addr == "" {
		log.Error("Address string is empty")
		return "", ErrEmptyAddress
	}
	output, err := utils.Execute("nslookup", addr)
	return string(output), err
}
